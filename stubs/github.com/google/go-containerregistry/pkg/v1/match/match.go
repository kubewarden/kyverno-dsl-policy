package match

import v1 "github.com/google/go-containerregistry/pkg/v1"

type Matcher func(desc v1.Descriptor) bool

func Name(name string) Matcher {
	panic("stub")
}

func Annotation(key, value string) Matcher {
	panic("stub")
}

func Platforms(platforms ...v1.Platform) Matcher {
	panic("stub")
}

func MediaTypes(mediaTypes ...string) Matcher {
	panic("stub")
}

func Digests(digests ...v1.Hash) Matcher {
	panic("stub")
}

type Embedme interface{}
