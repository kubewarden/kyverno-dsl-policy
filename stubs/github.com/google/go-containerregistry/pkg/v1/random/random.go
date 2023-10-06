package random

import (
	"io"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type uncompressedLayer struct {
	diffID    v1.Hash
	mediaType types.MediaType
	content   []byte
}

func (ul *uncompressedLayer) DiffID() (v1.Hash, error) {
	panic("stub")
}

func (ul *uncompressedLayer) Uncompressed() (io.ReadCloser, error) {
	panic("stub")
}

func (ul *uncompressedLayer) MediaType() (types.MediaType, error) {
	panic("stub")
}

func Image(byteSize, layers int64) (v1.Image, error) {
	panic("stub")
}

func Layer(byteSize int64, mt types.MediaType) (v1.Layer, error) {
	panic("stub")
}

type randomIndex struct {
	images   map[v1.Hash]v1.Image
	manifest *v1.IndexManifest
}

func Index(byteSize, layers, count int64) (v1.ImageIndex, error) {
	panic("stub")
}

func (i *randomIndex) MediaType() (types.MediaType, error) {
	panic("stub")
}

func (i *randomIndex) Digest() (v1.Hash, error) {
	panic("stub")
}

func (i *randomIndex) Size() (int64, error) {
	panic("stub")
}

func (i *randomIndex) IndexManifest() (*v1.IndexManifest, error) {
	panic("stub")
}

func (i *randomIndex) RawManifest() ([]byte, error) {
	panic("stub")
}

func (i *randomIndex) Image(h v1.Hash) (v1.Image, error) {
	panic("stub")
}

func (i *randomIndex) ImageIndex(h v1.Hash) (v1.ImageIndex, error) {
	panic("stub")
}

type Embedme interface{}
