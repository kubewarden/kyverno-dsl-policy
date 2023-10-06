package compression

import (
	"io"

	"github.com/google/go-containerregistry/pkg/compression"
)

type Opener func() (io.ReadCloser, error)

type PeekReader interface {
	Peek(n int) ([]byte, error)
}

func GetCompression(opener Opener) (compression.Compression, error) {
	panic("stub")
}

func PeekCompression(r io.Reader) (compression.Compression, PeekReader, error) {
	panic("stub")
}

type Embedme interface{}
