package layout

import (
	"io"
	"os"

	"sync"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/match"
	"github.com/google/go-containerregistry/pkg/v1/partial"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

func (l Path) Blob(h v1.Hash) (io.ReadCloser, error) {
	panic("stub")
}

func (l Path) Bytes(h v1.Hash) ([]byte, error) {
	panic("stub")
}

type compressedBlob struct {
	path Path
	desc v1.Descriptor
}

type layoutImage struct {
	path         Path
	desc         v1.Descriptor
	manifestLock sync.Mutex
	rawManifest  []byte
}

func (l Path) Image(h v1.Hash) (v1.Image, error) {
	panic("stub")
}

func (li *layoutImage) MediaType() (types.MediaType, error) {
	panic("stub")
}

func (li *layoutImage) Manifest() (*v1.Manifest, error) {
	panic("stub")
}

func (li *layoutImage) RawManifest() ([]byte, error) {
	panic("stub")
}

func (li *layoutImage) RawConfigFile() ([]byte, error) {
	panic("stub")
}

func (li *layoutImage) LayerByDigest(h v1.Hash) (partial.CompressedLayer, error) {
	panic("stub")
}

func (b *compressedBlob) Digest() (v1.Hash, error) {
	panic("stub")
}

func (b *compressedBlob) Compressed() (io.ReadCloser, error) {
	panic("stub")
}

func (b *compressedBlob) Size() (int64, error) {
	panic("stub")
}

func (b *compressedBlob) MediaType() (types.MediaType, error) {
	panic("stub")
}

func (b *compressedBlob) Descriptor() (*v1.Descriptor, error) {
	panic("stub")
}

func (b *compressedBlob) Exists() (bool, error) {
	panic("stub")
}

type layoutIndex struct {
	mediaType types.MediaType
	path      Path
	rawIndex  []byte
}

func ImageIndexFromPath(path string) (v1.ImageIndex, error) {
	panic("stub")
}

func (l Path) ImageIndex() (v1.ImageIndex, error) {
	panic("stub")
}

func (i *layoutIndex) MediaType() (types.MediaType, error) {
	panic("stub")
}

func (i *layoutIndex) Digest() (v1.Hash, error) {
	panic("stub")
}

func (i *layoutIndex) Size() (int64, error) {
	panic("stub")
}

func (i *layoutIndex) IndexManifest() (*v1.IndexManifest, error) {
	panic("stub")
}

func (i *layoutIndex) RawManifest() ([]byte, error) {
	panic("stub")
}

func (i *layoutIndex) Image(h v1.Hash) (v1.Image, error) {
	panic("stub")
}

func (i *layoutIndex) ImageIndex(h v1.Hash) (v1.ImageIndex, error) {
	panic("stub")
}

func (i *layoutIndex) Blob(h v1.Hash) (io.ReadCloser, error) {
	panic("stub")
}

type Path string

type Option func(*options)

type options struct{ descOpts []descriptorOption }

type descriptorOption func(*v1.Descriptor)

func WithAnnotations(annotations map[string]string) Option {
	panic("stub")
}

func WithURLs(urls []string) Option {
	panic("stub")
}

func WithPlatform(platform v1.Platform) Option {
	panic("stub")
}

func FromPath(path string) (Path, error) {
	panic("stub")
}

var layoutFile = `{
    "imageLayoutVersion": "1.0.0"
}`

type withLayer interface {
	Layer(v1.Hash) (v1.Layer, error)
}

type withBlob interface {
	Blob(v1.Hash) (io.ReadCloser, error)
}

func (l Path) AppendImage(img v1.Image, options ...Option) error {
	panic("stub")
}

func (l Path) AppendIndex(ii v1.ImageIndex, options ...Option) error {
	panic("stub")
}

func (l Path) AppendDescriptor(desc v1.Descriptor) error {
	panic("stub")
}

func (l Path) ReplaceImage(img v1.Image, matcher match.Matcher, options ...Option) error {
	panic("stub")
}

func (l Path) ReplaceIndex(ii v1.ImageIndex, matcher match.Matcher, options ...Option) error {
	panic("stub")
}

func (l Path) RemoveDescriptors(matcher match.Matcher) error {
	panic("stub")
}

func (l Path) WriteFile(name string, data []byte, perm os.FileMode) error {
	panic("stub")
}

func (l Path) WriteBlob(hash v1.Hash, r io.ReadCloser) error {
	panic("stub")
}

func (l Path) RemoveBlob(hash v1.Hash) error {
	panic("stub")
}

func (l Path) WriteImage(img v1.Image) error {
	panic("stub")
}

func (l Path) WriteIndex(ii v1.ImageIndex) error {
	panic("stub")
}

func Write(path string, ii v1.ImageIndex) (Path, error) {
	panic("stub")
}

type Embedme interface{}
