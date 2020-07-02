package httpclient

import "net/http"

type IServiceCall interface {
	GetRequest(r *http.Request)
	GetResponse() ([]byte, int)
}
