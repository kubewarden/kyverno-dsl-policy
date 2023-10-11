package retry

import (
	"context"

	"github.com/google/go-containerregistry/internal/retry/wait"
)

type temporary interface {
	Temporary() bool
}

type Predicate func(error) (retry bool)

type Backoff wait.Backoff

type contextKey string

func IsTemporary(err error) bool {
	panic("stub")
}

func IsNotNil(err error) bool {
	panic("stub")
}

func Retry(f func() error, p Predicate, backoff wait.Backoff) (err error) {
	panic("stub")
}

func Never(ctx context.Context) context.Context {
	panic("stub")
}

func Ever(ctx context.Context) bool {
	panic("stub")
}

type Embedme interface{}
