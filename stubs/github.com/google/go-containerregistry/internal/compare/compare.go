package compare

import v1 "github.com/google/go-containerregistry/pkg/v1"

func Images(a, b v1.Image) error {
	panic("stub")
}

func Indexes(a, b v1.ImageIndex) error {
	panic("stub")
}

func Layers(a, b v1.Layer) error {
	panic("stub")
}

type Embedme interface{}
