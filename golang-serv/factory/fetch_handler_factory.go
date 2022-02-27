package factory

import (
	"fmt"
	"net/http"
)

type FetchFactory struct{}

func (f *FetchFactory) All() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is %v\n", r.URL)
	}
}
