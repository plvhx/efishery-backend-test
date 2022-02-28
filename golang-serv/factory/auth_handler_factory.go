package factory

import (
	"fmt"
	"service/cache"
	"service/hash"
	"service/jwt"
	"service/password"
	"time"

	coreHttp "net/http"
	servHttp "service/http"
	request "service/request/auth"
	response "service/response/auth"

	"github.com/go-playground/validator/v10"
)

var (
	Validator = validator.New()
)

type AuthFactory struct{}

func (a *AuthFactory) Register() func(coreHttp.ResponseWriter, *coreHttp.Request) {
	return func(w coreHttp.ResponseWriter, r *coreHttp.Request) {
		if r.Method != "POST" {
			servHttp.HandleError(
				w,
				"HTTP method must be 'POST'.",
				coreHttp.StatusMethodNotAllowed,
			)

			return
		}

		if !servHttp.IsJson(r) {
			servHttp.HandleError(
				w,
				"Content-Type must be 'application/json'.",
				coreHttp.StatusUnsupportedMediaType,
			)

			return
		}

		req := request.NewRegister()
		err := servHttp.GetJson(r, &req)

		if err != nil {
			servHttp.HandleError(w, err.Error(), coreHttp.StatusUnprocessableEntity)
			return
		}

		err = Validator.Struct(req)

		if err != nil {
			servHttp.HandleError(w, err.Error(), coreHttp.StatusUnprocessableEntity)
			return
		}

		auth := response.NewAuth()

		auth.SetName(req.GetName())
		auth.SetPhone(req.GetPhone())
		auth.SetRole(req.GetRole())
		auth.SetPassword(password.Generate(4))

		c, err := cache.AuthCacheFactory()

		if err != nil {
			servHttp.HandleError(w, err.Error(), coreHttp.StatusInternalServerError)
			return
		}

		ckey := hash.Calculate(auth.GetName() + auth.GetPhone() + auth.GetRole())

		c.Put(ckey, auth)
		servHttp.AsJson(w, auth, coreHttp.StatusCreated)
	}
}

func (a *AuthFactory) Token() func(coreHttp.ResponseWriter, *coreHttp.Request) {
	return func(w coreHttp.ResponseWriter, r *coreHttp.Request) {
		if r.Method != "POST" {
			servHttp.HandleError(
				w,
				"HTTP method must be 'POST'.",
				coreHttp.StatusMethodNotAllowed,
			)

			return
		}

		if !servHttp.IsJson(r) {
			servHttp.HandleError(
				w,
				"Content-Type must be 'application/json'.",
				coreHttp.StatusUnsupportedMediaType,
			)

			return
		}

		req := request.NewToken()
		err := servHttp.GetJson(r, &req)

		if err != nil {
			servHttp.HandleError(w, err.Error(), coreHttp.StatusUnprocessableEntity)
			return
		}

		err = Validator.Struct(req)

		if err != nil {
			servHttp.HandleError(w, err.Error(), coreHttp.StatusUnprocessableEntity)
			return
		}

		c, err := cache.AuthCacheFactory()

		if err != nil {
			servHttp.HandleError(w, err.Error(), coreHttp.StatusInternalServerError)
			return
		}

		data := c.All()
		auth := response.NewAuth()

		var tval map[string]interface{}

		for _, abstract := range data.Data {
			concrete := abstract.(map[string]interface{})

			if concrete["phone"] == req.GetPhone() && concrete["password"] == req.GetPassword() {
				tval = concrete
				break
			}
		}

		if tval == nil {
			servHttp.HandleError(
				w,
				fmt.Sprintf(
					"Data not found (phone: %v, password: %v)",
					req.GetPhone(),
					req.GetPassword(),
				),
				coreHttp.StatusNotFound,
			)

			return
		}

		rval := make(map[string]string, 0)

		for key, value := range tval {
			converted, ok := value.(string)

			if !ok {
				servHttp.HandleError(
					w,
					"Cannot convert interface{} to string.",
					coreHttp.StatusInternalServerError,
				)

				return
			}

			rval[key] = converted
		}

		auth.SetName(rval["name"])
		auth.SetPhone(rval["phone"])
		auth.SetRole(rval["role"])
		auth.SetPassword(rval["password"])
		auth.SetTimestamp(time.Now().Unix())

		buf, err := jwt.Generate(auth)

		if err != nil {
			servHttp.HandleError(w, err.Error(), coreHttp.StatusInternalServerError)
			return
		}

		token := response.NewToken()
		token.SetToken(buf)

		servHttp.AsJson(w, token, coreHttp.StatusOK)
	}
}

func (a *AuthFactory) Parse() func(coreHttp.ResponseWriter, *coreHttp.Request) {
	return func(w coreHttp.ResponseWriter, r *coreHttp.Request) {
		if r.Method != "POST" {
			servHttp.HandleError(
				w,
				"HTTP method must be 'POST'.",
				coreHttp.StatusMethodNotAllowed,
			)

			return
		}

		if !servHttp.IsJson(r) {
			servHttp.HandleError(
				w,
				"Content-Type must be 'application/json'.",
				coreHttp.StatusUnsupportedMediaType,
			)

			return
		}

		req := request.NewParse()
		err := servHttp.GetJson(r, &req)

		if err != nil {
			servHttp.HandleError(w, err.Error(), coreHttp.StatusUnprocessableEntity)
			return
		}

		err = Validator.Struct(req)

		if err != nil {
			servHttp.HandleError(w, err.Error(), coreHttp.StatusUnprocessableEntity)
			return
		}

		auth, err := jwt.Parse(req.GetToken())

		if err != nil {
			servHttp.HandleError(w, err.Error(), coreHttp.StatusBadRequest)
			return
		}

		servHttp.AsJson(w, auth, coreHttp.StatusOK)
	}
}
