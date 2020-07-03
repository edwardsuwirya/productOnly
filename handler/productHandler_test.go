package handler

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewProductHandler(t *testing.T) {
	productHandlerMock := NewProductHandler()
	assert.NotNil(t, productHandlerMock)
}

func TestProductHandler_Handler(t *testing.T) {
	req, err := http.NewRequest("GET", "/product", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	productHandlerMock := NewProductHandler()
	handler := http.HandlerFunc(productHandlerMock.Handler)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, 200)

	assert.Equal(t, rr.Body.String(), "product")
}
