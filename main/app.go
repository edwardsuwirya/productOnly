package main

import (
	"github.com/edwardsuwirya/productOnly/handler"
	"github.com/edwardsuwirya/productOnly/middleware"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Use(middleware.ActivityLogMiddleware)

	prod := r.PathPrefix("/product").Subrouter()
	prod.Use(middleware.TokenValidationMiddleware)
	prod.HandleFunc("", handler.NewProductHandler().Handler).Methods(http.MethodGet)

	log.Print("Server is listening")
	if err := http.ListenAndServe("localhost:7001", r); err != nil {
		log.Panic(err)
	}
}
