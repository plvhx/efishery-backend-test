package factory

import (
	"service/jwt"
	"service/resource"

	coreHttp "net/http"
	servHttp "service/http"
)

type FetchFactory struct{}

func (f *FetchFactory) All() func(coreHttp.ResponseWriter, *coreHttp.Request) {
	return func(w coreHttp.ResponseWriter, r *coreHttp.Request) {
		if !servHttp.HasAuthorization(r) {
			servHttp.HandleError(
				w,
				"Authorization header must be provided.",
				coreHttp.StatusUnauthorized,
			)

			return
		}

		token, err := servHttp.ParseBearerToken(r)

		if err != nil {
			servHttp.HandleError(w, err.Error(), coreHttp.StatusUnauthorized)
			return
		}

		_, err = jwt.Parse(token)

		if err != nil {
			servHttp.HandleError(w, err.Error(), coreHttp.StatusUnauthorized)
			return
		}

		data, err := resource.Fetch()

		if err != nil {
			servHttp.HandleError(w, err.Error(), coreHttp.StatusInternalServerError)
			return
		}

		servHttp.AsJson(w, data, coreHttp.StatusOK)
	}
}
