package verify

import (
	"hash"
	"io"

	v1 "github.com/google/go-containerregistry/pkg/v1"
)

type verifyReader struct {
	inner    io.Reader
	hasher   hash.Hash
	expected v1.Hash
	gotSize
	wantSize int64
}

type Error struct {
	got     string
	want    v1.Hash
	gotSize int64
}

func (v Error) Error() string {
	panic("stub")
}

func (vc *verifyReader) Read(b []byte) (int, error) {
	panic("stub")
}

func ReadCloser(r io.ReadCloser, size int64, h v1.Hash) (io.ReadCloser, error) {
	panic("stub")
}

func Descriptor(d v1.Descriptor) error {
	panic("stub")
}

type Embedme interface{}
