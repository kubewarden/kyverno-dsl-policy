package stream

import (
	"io"
	"sync"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type Layer struct {
	blob        io.ReadCloser
	consumed    bool
	compression int
	mu          sync.Mutex
	digest
	diffID    *v1.Hash
	size      int64
	mediaType types.MediaType
}

type compressedReader struct {
	pr     io.Reader
	closer func() error
}

type countWriter struct{ n int64 }

type LayerOption func(*Layer)

func WithCompressionLevel(level int) LayerOption {
	panic("stub")
}

func WithMediaType(mt types.MediaType) LayerOption {
	panic("stub")
}

func NewLayer(rc io.ReadCloser, opts ...LayerOption) *Layer {
	panic("stub")
}

func (l *Layer) Digest() (v1.Hash, error) {
	panic("stub")
}

func (l *Layer) DiffID() (v1.Hash, error) {
	panic("stub")
}

func (l *Layer) Size() (int64, error) {
	panic("stub")
}

func (l *Layer) MediaType() (types.MediaType, error) {
	panic("stub")
}

func (l *Layer) Uncompressed() (io.ReadCloser, error) {
	panic("stub")
}

func (l *Layer) Compressed() (io.ReadCloser, error) {
	panic("stub")
}

func (cr *compressedReader) Read(b []byte) (int, error) {
	panic("stub")
}

func (cr *compressedReader) Close() error {
	panic("stub")
}

func (c *countWriter) Write(p []byte) (int, error) {
	panic("stub")
}

type Embedme interface{}
