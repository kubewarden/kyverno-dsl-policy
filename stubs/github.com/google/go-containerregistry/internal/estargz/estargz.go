package estargz

import (
	"io"

	v1 "github.com/google/go-containerregistry/pkg/v1"
)

func ReadCloser(r io.ReadCloser, opts ...interface{}) (interface{}, v1.Hash, error) {
	panic("stub")
}

type Embedme interface{}
