package redact

import (
	"context"
	"net/url"
)

type contextKey string

func NewContext(ctx context.Context, reason string) context.Context {
	panic("stub")
}

func FromContext(ctx context.Context) (bool, string) {
	panic("stub")
}

func Error(err error) error {
	panic("stub")
}

func URL(u *url.URL) *url.URL {
	panic("stub")
}

type Embedme interface{}
