package partial

import (
	"io"

	"sync"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/match"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type CompressedLayer interface {
	Digest() (v1.Hash, error)
	Compressed() (io.ReadCloser, error)
	Size() (int64, error)
	MediaType() (types.MediaType, error)
}

type compressedLayerExtender struct{ CompressedLayer }

type CompressedImageCore interface {
	RawManifest() ([]byte, error)
	LayerByDigest(v1.Hash) (CompressedLayer, error)
}

type compressedImageExtender struct{ CompressedImageCore }

func (cle *compressedLayerExtender) Uncompressed() (io.ReadCloser, error) {
	panic("stub")
}

func (cle *compressedLayerExtender) DiffID() (v1.Hash, error) {
	panic("stub")
}

func CompressedToLayer(ul CompressedLayer) (v1.Layer, error) {
	panic("stub")
}

func (i *compressedImageExtender) Digest() (v1.Hash, error) {
	panic("stub")
}

func (i *compressedImageExtender) ConfigName() (v1.Hash, error) {
	panic("stub")
}

func (i *compressedImageExtender) Layers() ([]v1.Layer, error) {
	panic("stub")
}

func (i *compressedImageExtender) LayerByDigest(h v1.Hash) (v1.Layer, error) {
	panic("stub")
}

func (i *compressedImageExtender) LayerByDiffID(h v1.Hash) (v1.Layer, error) {
	panic("stub")
}

func (i *compressedImageExtender) ConfigFile() (*v1.ConfigFile, error) {
	panic("stub")
}

func (i *compressedImageExtender) Manifest() (*v1.Manifest, error) {
	panic("stub")
}

func (i *compressedImageExtender) Size() (int64, error) {
	panic("stub")
}

func CompressedToImage(cic CompressedImageCore) (v1.Image, error) {
	panic("stub")
}

type ImageCore interface {
	RawConfigFile() ([]byte, error)
	MediaType() (types.MediaType, error)
}

func FindManifests(index v1.ImageIndex, matcher match.Matcher) ([]v1.Descriptor, error) {
	panic("stub")
}

func FindImages(index v1.ImageIndex, matcher match.Matcher) ([]v1.Image, error) {
	panic("stub")
}

func FindIndexes(index v1.ImageIndex, matcher match.Matcher) ([]v1.ImageIndex, error) {
	panic("stub")
}

type uncompressedImageExtender struct {
	UncompressedImageCore
	lock     sync.Mutex
	manifest *v1.Manifest
}

type UncompressedLayer interface {
	DiffID() (v1.Hash, error)
	Uncompressed() (io.ReadCloser, error)
	MediaType() (types.MediaType, error)
}

type uncompressedLayerExtender struct {
	UncompressedLayer
	hash          v1.Hash
	size          int64
	hashSizeError error
	once          sync.Once
}

type UncompressedImageCore interface {
	LayerByDiffID(v1.Hash) (UncompressedLayer, error)
}

func (ule *uncompressedLayerExtender) Compressed() (io.ReadCloser, error) {
	panic("stub")
}

func (ule *uncompressedLayerExtender) Digest() (v1.Hash, error) {
	panic("stub")
}

func (ule *uncompressedLayerExtender) Size() (int64, error) {
	panic("stub")
}

func UncompressedToLayer(ul UncompressedLayer) (v1.Layer, error) {
	panic("stub")
}

func UncompressedToImage(uic UncompressedImageCore) (v1.Image, error) {
	panic("stub")
}

func (i *uncompressedImageExtender) Digest() (v1.Hash, error) {
	panic("stub")
}

func (i *uncompressedImageExtender) Manifest() (*v1.Manifest, error) {
	panic("stub")
}

func (i *uncompressedImageExtender) RawManifest() ([]byte, error) {
	panic("stub")
}

func (i *uncompressedImageExtender) Size() (int64, error) {
	panic("stub")
}

func (i *uncompressedImageExtender) ConfigName() (v1.Hash, error) {
	panic("stub")
}

func (i *uncompressedImageExtender) ConfigFile() (*v1.ConfigFile, error) {
	panic("stub")
}

func (i *uncompressedImageExtender) Layers() ([]v1.Layer, error) {
	panic("stub")
}

func (i *uncompressedImageExtender) LayerByDiffID(diffID v1.Hash) (v1.Layer, error) {
	panic("stub")
}

func (i *uncompressedImageExtender) LayerByDigest(h v1.Hash) (v1.Layer, error) {
	panic("stub")
}

type WithRawManifest interface {
	RawManifest() ([]byte, error)
}

type WithDiffID interface {
	DiffID() (v1.Hash, error)
}

type withArtifactType interface {
	ArtifactType() (string, error)
}

type withDescriptor interface {
	Descriptor() (*v1.Descriptor, error)
}

type WithManifest interface {
	Manifest() (*v1.Manifest, error)
}

type withExists interface {
	Exists() (bool, error)
}

type withUncompressedSize interface {
	UncompressedSize() (int64, error)
}

type WithRawConfigFile interface {
	RawConfigFile() ([]byte, error)
}

type withConfigLayer interface {
	ConfigLayer() (v1.Layer, error)
}

type WithConfigFile interface {
	ConfigFile() (*v1.ConfigFile, error)
}

type WithManifestAndConfigFile interface {
	Manifest() (*v1.Manifest, error)
}

type Describable interface {
	Digest() (v1.Hash, error)
	MediaType() (types.MediaType, error)
	Size() (int64, error)
}

type configLayer struct {
	hash    v1.Hash
	content []byte
}

func ConfigFile(i WithRawConfigFile) (*v1.ConfigFile, error) {
	panic("stub")
}

func ConfigName(i WithRawConfigFile) (v1.Hash, error) {
	panic("stub")
}

func (cl *configLayer) Digest() (v1.Hash, error) {
	panic("stub")
}

func (cl *configLayer) DiffID() (v1.Hash, error) {
	panic("stub")
}

func (cl *configLayer) Uncompressed() (io.ReadCloser, error) {
	panic("stub")
}

func (cl *configLayer) Compressed() (io.ReadCloser, error) {
	panic("stub")
}

func (cl *configLayer) Size() (int64, error) {
	panic("stub")
}

func (cl *configLayer) MediaType() (types.MediaType, error) {
	panic("stub")
}

func ConfigLayer(i WithRawConfigFile) (v1.Layer, error) {
	panic("stub")
}

func DiffIDs(i WithConfigFile) ([]v1.Hash, error) {
	panic("stub")
}

func RawConfigFile(i WithConfigFile) ([]byte, error) {
	panic("stub")
}

func Digest(i WithRawManifest) (v1.Hash, error) {
	panic("stub")
}

func Manifest(i WithRawManifest) (*v1.Manifest, error) {
	panic("stub")
}

func RawManifest(i WithManifest) ([]byte, error) {
	panic("stub")
}

func Size(i WithRawManifest) (int64, error) {
	panic("stub")
}

func FSLayers(i WithManifest) ([]v1.Hash, error) {
	panic("stub")
}

func BlobSize(i WithManifest, h v1.Hash) (int64, error) {
	panic("stub")
}

func BlobDescriptor(i WithManifest, h v1.Hash) (*v1.Descriptor, error) {
	panic("stub")
}

func BlobToDiffID(i WithManifestAndConfigFile, h v1.Hash) (v1.Hash, error) {
	panic("stub")
}

func DiffIDToBlob(wm WithManifestAndConfigFile, h v1.Hash) (v1.Hash, error) {
	panic("stub")
}

func Descriptor(d Describable) (*v1.Descriptor, error) {
	panic("stub")
}

func UncompressedSize(l v1.Layer) (int64, error) {
	panic("stub")
}

func Exists(l v1.Layer) (bool, error) {
	panic("stub")
}

func ArtifactType(w WithManifest) (string, error) {
	panic("stub")
}

type Embedme interface{}
