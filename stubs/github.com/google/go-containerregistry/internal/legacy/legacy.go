package legacy

import (
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

type fslayer struct{ BlobSum string }

type schema1 struct{ FSLayers []fslayer }

func CopySchema1(desc *remote.Descriptor, srcRef, dstRef name.Reference, opts ...remote.Option) error {
	panic("stub")
}

type Embedme interface{}
