package tarball

import (
	"io"

	"github.com/google/go-containerregistry/pkg/legacy"
	"github.com/google/go-containerregistry/pkg/name"

	v1 "github.com/google/go-containerregistry/pkg/v1"
)

type repositoriesTarDescriptor map[string]map[string]string

type v1Layer struct {
	config *legacy.LayerConfigFile
	layer  v1.Layer
}

func Write(ref name.Reference, img v1.Image, w io.Writer) error {
	panic("stub")
}

func MultiWrite(refToImage map[name.Reference]v1.Image, w io.Writer) error {
	panic("stub")
}

type Embedme interface{}
