package httpclient

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewTokenServiceCall(t *testing.T) {
	t.Run("It should return Token Service", func(t *testing.T) {
		tokenServiceCall, err := NewTokenServiceCall("MOCK_URL")
		assert.Nil(t, err)
		assert.NotNil(t, tokenServiceCall)
	})
}

func TestTokenServiceCall_GetRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/product", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Authorization", "Bearer MOCK")
	tokenService, _ := NewTokenServiceCall("MOCK_URL")
	tokenService.GetRequest(req)
	assert.Equal(t, (tokenService.(*tokenServiceCall)).token, "MOCK")
}

func TestTokenServiceCall_GetResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("OK"))
	}))
	defer server.Close()

	tokenService, _ := NewTokenServiceCall(server.URL)
	tokenService.(*tokenServiceCall).token = "MOCK_TOKEN"
	body, responseCode := tokenService.GetResponse()
	assert.Equal(t, []byte("OK"), body)
	assert.Equal(t, 200, responseCode)
}
