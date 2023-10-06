package tarball

import (
	"io"
	"sync"

	"github.com/google/go-containerregistry/pkg/compression"
	"github.com/google/go-containerregistry/pkg/name"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/partial"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type uncompressedImage struct{ *image }

type Opener func() (io.ReadCloser, error)

type compressedLayerFromTarball struct {
	desc     v1.Descriptor
	opener   Opener
	filePath string
}

type tarFile struct {
	io.Reader
	io.Closer
}

type uncompressedLayerFromTarball struct {
	diffID    v1.Hash
	mediaType types.MediaType
	opener    Opener
	filePath  string
}

type compressedImage struct {
	*image
	manifestLock sync.Mutex
	manifest     *v1.Manifest
}

type Manifest []Descriptor

type foreignUncompressedLayer struct {
	uncompressedLayerFromTarball
	desc v1.Descriptor
}

type image struct {
	opener        Opener
	manifest      *Manifest
	config        []byte
	imgDescriptor *Descriptor
	tag           *name.Tag
}

type Descriptor struct {
	Config       string
	RepoTags     []string
	Layers       []string
	LayerSources map[v1.Hash]v1.Descriptor
}

func ImageFromPath(path string, tag *name.Tag) (v1.Image, error) {
	panic("stub")
}

func LoadManifest(opener Opener) (Manifest, error) {
	panic("stub")
}

func Image(opener Opener, tag *name.Tag) (v1.Image, error) {
	panic("stub")
}

func (i *image) MediaType() (types.MediaType, error) {
	panic("stub")
}

func (i *image) RawConfigFile() ([]byte, error) {
	panic("stub")
}

func (fl *foreignUncompressedLayer) Descriptor() (*v1.Descriptor, error) {
	panic("stub")
}

func (ulft *uncompressedLayerFromTarball) DiffID() (v1.Hash, error) {
	panic("stub")
}

func (ulft *uncompressedLayerFromTarball) Uncompressed() (io.ReadCloser, error) {
	panic("stub")
}

func (ulft *uncompressedLayerFromTarball) MediaType() (types.MediaType, error) {
	panic("stub")
}

func (i *uncompressedImage) LayerByDiffID(h v1.Hash) (partial.UncompressedLayer, error) {
	panic("stub")
}

func (c *compressedImage) Manifest() (*v1.Manifest, error) {
	panic("stub")
}

func (c *compressedImage) RawManifest() ([]byte, error) {
	panic("stub")
}

func (clft *compressedLayerFromTarball) Digest() (v1.Hash, error) {
	panic("stub")
}

func (clft *compressedLayerFromTarball) Compressed() (io.ReadCloser, error) {
	panic("stub")
}

func (clft *compressedLayerFromTarball) MediaType() (types.MediaType, error) {
	panic("stub")
}

func (clft *compressedLayerFromTarball) Size() (int64, error) {
	panic("stub")
}

func (c *compressedImage) LayerByDigest(h v1.Hash) (partial.CompressedLayer, error) {
	panic("stub")
}

type LayerOption func(*layer)

type layer struct {
	digest             v1.Hash
	diffID             v1.Hash
	size               int64
	compressedopener   Opener
	uncompressedopener Opener
	compression        compression.Compression
	compressionLevel   int
	annotations        map[string]string
	estgzopts          []interface{}
	mediaType          types.MediaType
}

func (l *layer) Descriptor() (*v1.Descriptor, error) {
	panic("stub")
}

func (l *layer) Digest() (v1.Hash, error) {
	panic("stub")
}

func (l *layer) DiffID() (v1.Hash, error) {
	panic("stub")
}

func (l *layer) Compressed() (io.ReadCloser, error) {
	panic("stub")
}

func (l *layer) Uncompressed() (io.ReadCloser, error) {
	panic("stub")
}

func (l *layer) Size() (int64, error) {
	panic("stub")
}

func (l *layer) MediaType() (types.MediaType, error) {
	panic("stub")
}

func WithCompression(comp compression.Compression) LayerOption {
	panic("stub")
}

func WithCompressionLevel(level int) LayerOption {
	panic("stub")
}

func WithMediaType(mt types.MediaType) LayerOption {
	panic("stub")
}

func WithCompressedCaching(l *layer) {
	panic("stub")
}

func WithEstargzOptions(opts ...interface{}) LayerOption {
	panic("stub")
}

func WithEstargz(l *layer) {
	panic("stub")
}

func LayerFromFile(path string, opts ...LayerOption) (v1.Layer, error) {
	panic("stub")
}

func LayerFromOpener(opener Opener, opts ...LayerOption) (v1.Layer, error) {
	panic("stub")
}

func LayerFromReader(reader io.Reader, opts ...LayerOption) (v1.Layer, error) {
	panic("stub")
}

type WriteOption func(*writeOptions) error

type progressWriter struct {
	w       io.Writer
	updates chan<- v1.Update
	size
	complete int64
}

type writeOptions struct{ updates chan<- v1.Update }

func WriteToFile(p string, ref name.Reference, img v1.Image, opts ...WriteOption) error {
	panic("stub")
}

func MultiWriteToFile(p string, tagToImage map[name.Tag]v1.Image, opts ...WriteOption) error {
	panic("stub")
}

func MultiRefWriteToFile(p string, refToImage map[name.Reference]v1.Image, opts ...WriteOption) error {
	panic("stub")
}

func Write(ref name.Reference, img v1.Image, w io.Writer, opts ...WriteOption) error {
	panic("stub")
}

func MultiWrite(tagToImage map[name.Tag]v1.Image, w io.Writer, opts ...WriteOption) error {
	panic("stub")
}

func MultiRefWrite(refToImage map[name.Reference]v1.Image, w io.Writer, opts ...WriteOption) error {
	panic("stub")
}

func CalculateSize(refToImage map[name.Reference]v1.Image) (size int64, err error) {
	panic("stub")
}

func ComputeManifest(refToImage map[name.Reference]v1.Image) (Manifest, error) {
	panic("stub")
}

func WithProgress(updates chan<- v1.Update) WriteOption {
	panic("stub")
}

func (pw *progressWriter) Write(p []byte) (int, error) {
	panic("stub")
}

func (pw *progressWriter) Error(err error) error {
	panic("stub")
}

func (pw *progressWriter) Close() error {
	panic("stub")
}

type Embedme interface{}
