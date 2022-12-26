package main

import (
	"github.com/gorilla/mux"
	"github.com/toghrul-nasirli/go-bookstore/pkg/routes"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9010", router))
}
