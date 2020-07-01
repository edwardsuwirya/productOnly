package handler

import "net/http"

type IHandler interface {
	Handler(http.ResponseWriter, *http.Request)
}
