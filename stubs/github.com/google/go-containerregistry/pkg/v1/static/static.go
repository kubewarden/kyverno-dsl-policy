package static

import (
	"io"
	"sync"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type staticLayer struct {
	b    []byte
	mt   types.MediaType
	once sync.Once
	h    v1.Hash
}

func NewLayer(b []byte, mt types.MediaType) v1.Layer {
	panic("stub")
}

func (l *staticLayer) Digest() (v1.Hash, error) {
	panic("stub")
}

func (l *staticLayer) DiffID() (v1.Hash, error) {
	panic("stub")
}

func (l *staticLayer) Compressed() (io.ReadCloser, error) {
	panic("stub")
}

func (l *staticLayer) Uncompressed() (io.ReadCloser, error) {
	panic("stub")
}

func (l *staticLayer) Size() (int64, error) {
	panic("stub")
}

func (l *staticLayer) MediaType() (types.MediaType, error) {
	panic("stub")
}

type Embedme interface{}
