package middleware

import "net/http"

type IMiddleware interface {
	Intercept(next http.Handler) http.Handler
}
