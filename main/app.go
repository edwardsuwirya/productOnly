package main

import (
	"github.com/edwardsuwirya/productOnly/handler"
	"github.com/edwardsuwirya/productOnly/httpclient"
	"github.com/edwardsuwirya/productOnly/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	const AUTH_URL = "http://localhost:7000/auth/tokenValidation"
	tokenService, err := httpclient.NewTokenServiceCall(AUTH_URL)
	if err == nil {
		r := mux.NewRouter()
		r.Use(middleware.ActivityLogMiddleware)
		prod := r.PathPrefix("/product").Subrouter()
		prod.Use(middleware.NewTokenValidationMiddleware(tokenService).Intercept)
		prod.HandleFunc("", handler.NewProductHandler().Handler).Methods(http.MethodGet)

		log.Print("Server is listening")
		if err := http.ListenAndServe("localhost:7001", r); err != nil {
			log.Panic(err)
		}
	} else {
		panic("Token Service is down")
	}

}
