package main

import (
	"log"
	"net/http"

	"github.com/golang-API/endpoint"
	"github.com/golang-API/example"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	exampleMethod(r)
	biodataEndpoint(r)
	notFound(r)
	log.Fatal(http.ListenAndServe(":8181", r))
}

func exampleMethod(r *mux.Router) error {
	r.HandleFunc("/get", example.Get).Methods(http.MethodGet)
	r.HandleFunc("/post", example.Post).Methods(http.MethodPost)
	r.HandleFunc("/put", example.Put).Methods(http.MethodPut)
	r.HandleFunc("/patch", example.Patch).Methods(http.MethodPatch)
	r.HandleFunc("/delete", example.Delete).Methods(http.MethodDelete)

	return nil
}

func biodataEndpoint(r *mux.Router) error {
	r.HandleFunc("/biodata", endpoint.GetAllBiodata).Methods(http.MethodGet)
	r.HandleFunc("/biodata/{id}", endpoint.GetOneBiodata).Methods(http.MethodGet)

	return nil
}

func notFound(r *mux.Router) error {
	r.HandleFunc("/", endpoint.NotFound).Methods(http.MethodGet)

	return nil
}
