package middleware

import (
	"fmt"
	"github.com/edwardsuwirya/productOnly/httpclient"
	"net/http"
)

type tokenValidationMiddleware struct {
	service httpclient.IServiceCall
	h       http.Handler
}

func NewTokenValidationMiddleware(service httpclient.IServiceCall) IMiddleware {
	return &tokenValidationMiddleware{service, nil}
}

func (mdw *tokenValidationMiddleware) Intercept(next http.Handler) http.Handler {
	mdw.h = next
	return http.HandlerFunc(mdw.handler)
}

func (mdw *tokenValidationMiddleware) handler(w http.ResponseWriter, r *http.Request) {
	mdw.service.GetRequest(r)
	body, resp := mdw.service.GetResponse()

	defer func() {
		if err := recover(); err != nil {
			http.Error(w, "Service is Down", http.StatusServiceUnavailable)
		}
	}()

	switch resp {
	case 200:
		fmt.Println(string(body))
		mdw.h.ServeHTTP(w, r)
	case 500:
		http.Error(w, "Internal Server Error", resp)
	case 503:
		http.Error(w, "Service Down", resp)
	default:
		http.Error(w, "You are not allowed to access this service", http.StatusUnauthorized)
	}
}
