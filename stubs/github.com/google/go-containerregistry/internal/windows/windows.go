package windows

import v1 "github.com/google/go-containerregistry/pkg/v1"

const userOwnerAndGroupSID = "AQAAgBQAAAAkAAAAAAAAAAAAAAABAgAAAAAABSAAAAAhAgAAAQIAAAAAAAUgAAAAIQIAAA=="

func Windows(layer v1.Layer) (v1.Layer, error) {
	panic("stub")
}

type Embedme interface{}
