package middleware

import (
	"log"
	"net/http"
)

func ActivityLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Bebas berkarya
		userAgent := r.Header.Get("User-Agent")
		log.Printf("Accessing path %v with application %v ", r.RequestURI, userAgent)
		next.ServeHTTP(w, r)
	})
}
