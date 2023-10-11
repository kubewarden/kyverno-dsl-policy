package httptest

import (
	"net/http"
	"net/http/httptest"
)

func NewTLSServer(domain string, handler http.Handler) (*httptest.Server, error) {
	panic("stub")
}

type Embedme interface{}
