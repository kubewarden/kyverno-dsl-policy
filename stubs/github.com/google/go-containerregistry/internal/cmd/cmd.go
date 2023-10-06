package cmd

import (
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type withMediaType struct{ MediaType string }

type rawManifest struct {
	body      []byte
	mediaType types.MediaType
}

func NewCmdEdit(options *[]crane.Option) interface{} {
	panic("stub")
}

func NewCmdEditConfig(options *[]crane.Option) interface{} {
	panic("stub")
}

func NewCmdEditManifest(options *[]crane.Option) interface{} {
	panic("stub")
}

func NewCmdEditFs(options *[]crane.Option) interface{} {
	panic("stub")
}

func (r *rawManifest) RawManifest() ([]byte, error) {
	panic("stub")
}

func (r *rawManifest) MediaType() (types.MediaType, error) {
	panic("stub")
}

type Embedme interface{}
