package empty

import (
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/partial"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type emptyImage struct{}

func (i emptyImage) MediaType() (types.MediaType, error) {
	panic("stub")
}

func (i emptyImage) RawConfigFile() ([]byte, error) {
	panic("stub")
}

func (i emptyImage) ConfigFile() (*v1.ConfigFile, error) {
	panic("stub")
}

func (i emptyImage) LayerByDiffID(h v1.Hash) (partial.UncompressedLayer, error) {
	panic("stub")
}

type emptyIndex struct{}

func (i emptyIndex) MediaType() (types.MediaType, error) {
	panic("stub")
}

func (i emptyIndex) Digest() (v1.Hash, error) {
	panic("stub")
}

func (i emptyIndex) Size() (int64, error) {
	panic("stub")
}

func (i emptyIndex) IndexManifest() (*v1.IndexManifest, error) {
	panic("stub")
}

func (i emptyIndex) RawManifest() ([]byte, error) {
	panic("stub")
}

func (i emptyIndex) Image(v1.Hash) (v1.Image, error) {
	panic("stub")
}

func (i emptyIndex) ImageIndex(v1.Hash) (v1.ImageIndex, error) {
	panic("stub")
}

type Embedme interface{}
