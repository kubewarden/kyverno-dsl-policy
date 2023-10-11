package validate

import (
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

func Image(img v1.Image, opt ...Option) error {
	panic("stub")
}

type withLayer interface {
	Layer(v1.Hash) (v1.Layer, error)
}

type withMediaType interface {
	MediaType() (types.MediaType, error)
}

func Index(idx v1.ImageIndex, opt ...Option) error {
	panic("stub")
}

type computedLayer struct {
	digest             v1.Hash
	size               int64
	diffid             v1.Hash
	uncompressedDiffid v1.Hash
	uncompressedSize   int64
}

func Layer(layer v1.Layer, opt ...Option) error {
	panic("stub")
}

type Option func(*options)

type options struct{ fast bool }

func Fast(o *options) {
	panic("stub")
}

type Embedme interface{}
