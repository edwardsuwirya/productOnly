package middleware

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type tokenServiceMock struct{}

func (t tokenServiceMock) GetRequest(r *http.Request) {
}

func (t tokenServiceMock) GetResponse() ([]byte, int) {
	return []byte("ok"), 200
}

func nextHandlerMock(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func TestTokenValidationMiddleware_handler(t *testing.T) {
	t.Run("It should return 200", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/product", nil)
		if err != nil {
			t.Fatal(err)
		}

		mdwMock := NewTokenValidationMiddleware(tokenServiceMock{})

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(mdwMock.(*tokenValidationMiddleware).handler)
		mdwMock.(*tokenValidationMiddleware).h = http.HandlerFunc(nextHandlerMock)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, 200)
	})
}
