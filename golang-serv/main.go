package main

import (
	"log"
	"net/http"
	"service/factory"
)

var (
	auth  = &factory.AuthFactory{}
	fetch = &factory.FetchFactory{}
)

func main() {
	http.HandleFunc("/auth/register", auth.Register())
	http.HandleFunc("/auth/token", auth.Token())
	http.HandleFunc("/auth/parse", auth.Parse())
	http.HandleFunc("/fetch", fetch.All())
	log.Fatal(http.ListenAndServe(":9000", nil))
}
