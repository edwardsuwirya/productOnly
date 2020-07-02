package middleware

import (
	"fmt"
	"github.com/edwardsuwirya/productOnly/httpclient"
	"net/http"
)

type tokenValidationMiddleware struct {
	service httpclient.IServiceCall
}

func NewTokenValidationMiddleware(service httpclient.IServiceCall) IMiddleware {
	return &tokenValidationMiddleware{service}
}

func (mdw *tokenValidationMiddleware) Intercept(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mdw.service.GetRequest(r)
		body, resp := mdw.service.GetResponse()

		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Service Down", http.StatusServiceUnavailable)
			}
		}()

		switch resp {
		case 200:
			fmt.Println(string(body))
			next.ServeHTTP(w, r)
		case 500:
			http.Error(w, "Internal Server Error", resp)
		case 503:
			http.Error(w, "Service Down", resp)
		default:
			http.Error(w, "You are not allowed to access this service", http.StatusUnauthorized)
		}
	})
}
