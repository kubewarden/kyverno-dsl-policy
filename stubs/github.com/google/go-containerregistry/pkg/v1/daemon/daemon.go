package daemon

import (
	"context"
	"io"
	"sync"

	"github.com/google/go-containerregistry/pkg/name"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type image struct {
	ref          name.Reference
	opener       *imageOpener
	tarballImage v1.Image
	id           *v1.Hash
	once         sync.Once
	err          error
}

type imageOpener struct {
	ref      name.Reference
	ctx      context.Context
	buffered bool
	client   Client
	once     sync.Once
	bytes    []byte
	err      error
}

func Image(ref name.Reference, options ...Option) (v1.Image, error) {
	panic("stub")
}

func (i *image) Layers() ([]v1.Layer, error) {
	panic("stub")
}

func (i *image) MediaType() (types.MediaType, error) {
	panic("stub")
}

func (i *image) Size() (int64, error) {
	panic("stub")
}

func (i *image) ConfigName() (v1.Hash, error) {
	panic("stub")
}

func (i *image) ConfigFile() (*v1.ConfigFile, error) {
	panic("stub")
}

func (i *image) RawConfigFile() ([]byte, error) {
	panic("stub")
}

func (i *image) Digest() (v1.Hash, error) {
	panic("stub")
}

func (i *image) Manifest() (*v1.Manifest, error) {
	panic("stub")
}

func (i *image) RawManifest() ([]byte, error) {
	panic("stub")
}

func (i *image) LayerByDigest(h v1.Hash) (v1.Layer, error) {
	panic("stub")
}

func (i *image) LayerByDiffID(h v1.Hash) (v1.Layer, error) {
	panic("stub")
}

type Client interface {
	NegotiateAPIVersion(ctx context.Context)
	ImageSave(context.Context, []string) (io.ReadCloser, error)
	ImageLoad(context.Context, io.Reader, bool) (types.ImageLoadResponse, error)
	ImageTag(context.Context, string, string) error
	ImageInspectWithRaw(context.Context, string) (types.ImageInspect, []byte, error)
}

type ImageOption Option

type Option func(*options)

type options struct {
	ctx      context.Context
	client   Client
	buffered bool
}

func WithBufferedOpener() Option {
	panic("stub")
}

func WithUnbufferedOpener() Option {
	panic("stub")
}

func WithClient(client Client) Option {
	panic("stub")
}

func WithContext(ctx context.Context) Option {
	panic("stub")
}

func Tag(src, dest name.Tag, options ...Option) error {
	panic("stub")
}

func Write(tag name.Tag, img v1.Image, options ...Option) (string, error) {
	panic("stub")
}

type Embedme interface{}
