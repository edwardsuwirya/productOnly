package httpclient

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type tokenServiceCall struct {
	url   string
	token string
}

func NewTokenServiceCall(url string) (IServiceCall, error) {
	if len(url) == 0 {
		return nil, errors.New("No URL defined")
	}
	return &tokenServiceCall{url: url}, nil
}

func (s *tokenServiceCall) GetRequest(r *http.Request) {
	token := r.Header.Get("Authorization")
	if len(token) == 0 {
		s.token = ""
	} else {
		tokenVal := strings.Split(token, "Bearer ")
		s.token = strings.Trim(tokenVal[1], "")
	}
}

func (s *tokenServiceCall) GetResponse() ([]byte, int) {
	resp, err := http.Get(fmt.Sprintf("%s?token=%s", s.url, s.token))
	defer func() {
		if err := recover(); err != nil {
			log.Printf("%s Service Down", s.url)
			panic("error token service call")
		}
	}()
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return nil, http.StatusServiceUnavailable
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	return body, http.StatusOK
}
