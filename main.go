package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var port = "6969"

func main() {
	log.Fatal(http.ListenAndServe(":"+port, router()))
}

func router() http.Handler {
	r := mux.NewRouter()
	r.Path("/hi").Methods(http.MethodGet).HandlerFunc(greet)
	return r
}

func greet(w http.ResponseWriter, req *http.Request) {
	_, _ = w.Write([]byte("Go la la!"))
}
