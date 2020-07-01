package middleware

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const AUTH_URL = "http://128.199.157.102:7000/auth/tokenValidation"

func TokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println(r.Header.Get("MyCustomHeader"))

		token := r.Header.Get("Authorization")
		if len(token) == 0 {
			http.Error(w, "You are not allowed to access this service", http.StatusUnauthorized)
		} else {
			tokenVal := strings.Split(token, "Bearer ")
			token := strings.Trim(tokenVal[1], "")

			resp, err := http.Get(fmt.Sprintf("%s?token=%s", AUTH_URL, token))

			defer func() {
				if err := recover(); err != nil {
					log.Printf("%s Service Down", AUTH_URL)
					http.Error(w, "Service Down", http.StatusServiceUnavailable)
				}
			}()
			defer resp.Body.Close()
			if err != nil {
				fmt.Println(err)
				http.Error(w, "You are not allowed to access this service", http.StatusUnauthorized)
			}
			if resp.StatusCode == 200 {
				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Println(string(body))
				next.ServeHTTP(w, r)
			} else if resp.StatusCode == 502 {
				log.Printf("%s Service Down", AUTH_URL)
				http.Error(w, "Service Down", http.StatusServiceUnavailable)
			} else {
				http.Error(w, "You are not allowed to access this service", http.StatusUnauthorized)
			}
		}
	})
}
