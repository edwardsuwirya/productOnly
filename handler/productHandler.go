package handler

import (
	"net/http"
)

type productHandler struct {
}

func NewProductHandler() IHandler {
	return &productHandler{}
}

func (h *productHandler) Handler(w http.ResponseWriter, r *http.Request) {
	//tanpa middleware
	//log.Printf("Accessing : %v", r.RequestURI)
	w.Write([]byte("product"))
}
