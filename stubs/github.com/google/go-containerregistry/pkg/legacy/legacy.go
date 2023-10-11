package legacy

import v1 "github.com/google/go-containerregistry/pkg/v1"

type LayerConfigFile struct {
	v1.ConfigFile
	ContainerConfig v1.Config
	ID              string
	Parent          string
	Throwaway       bool
	Comment         string
}

type Embedme interface{}
