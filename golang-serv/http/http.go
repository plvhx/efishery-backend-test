package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type FaultResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func IsJson(r *http.Request) bool {
	if r.Header.Get("Content-Type") != "application/json" {
		return false
	}

	return true
}

func AsJson(w http.ResponseWriter, v interface{}, code int) {
	buf, err := json.Marshal(v)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(buf)
}

func GetJson(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, &v)

	if err != nil {
		return err
	}

	return nil
}

func HandleError(w http.ResponseWriter, message string, code int) {
	AsJson(w, &FaultResponse{Message: message, Code: code}, code)
}
