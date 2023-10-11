package mutate

import (
	"io"
	"time"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/match"
	"github.com/google/go-containerregistry/pkg/v1/partial"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type image struct {
	base            v1.Image
	adds            []Addendum
	computed        bool
	configFile      *v1.ConfigFile
	manifest        *v1.Manifest
	annotations     map[string]string
	mediaType       *types.MediaType
	configMediaType *types.MediaType
	diffIDMap       map[v1.Hash]v1.Layer
	digestMap       map[v1.Hash]v1.Layer
	subject         *v1.Descriptor
}

func (i *image) MediaType() (types.MediaType, error) {
	panic("stub")
}

func (i *image) Layers() ([]v1.Layer, error) {
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

func (i *image) Size() (int64, error) {
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

type withLayer interface {
	Layer(v1.Hash) (v1.Layer, error)
}

type index struct {
	base        v1.ImageIndex
	adds        []IndexAddendum
	remove      match.Matcher
	computed    bool
	manifest    *v1.IndexManifest
	annotations map[string]string
	mediaType   *types.MediaType
	imageMap    map[v1.Hash]v1.Image
	indexMap    map[v1.Hash]v1.ImageIndex
	layerMap    map[v1.Hash]v1.Layer
	subject     *v1.Descriptor
}

func (i *index) MediaType() (types.MediaType, error) {
	panic("stub")
}

func (i *index) Size() (int64, error) {
	panic("stub")
}

func (i *index) Image(h v1.Hash) (v1.Image, error) {
	panic("stub")
}

func (i *index) ImageIndex(h v1.Hash) (v1.ImageIndex, error) {
	panic("stub")
}

func (i *index) Layer(h v1.Hash) (v1.Layer, error) {
	panic("stub")
}

func (i *index) Digest() (v1.Hash, error) {
	panic("stub")
}

func (i *index) IndexManifest() (*v1.IndexManifest, error) {
	panic("stub")
}

func (i *index) RawManifest() ([]byte, error) {
	panic("stub")
}

const whiteoutPrefix = ".wh."

type arbitraryRawManifest struct {
	a       partial.WithRawManifest
	anns    map[string]string
	subject *v1.Descriptor
}

type Addendum struct {
	Layer       v1.Layer
	History     v1.History
	URLs        []string
	Annotations map[string]string
	MediaType   types.MediaType
}

type Appendable interface {
	MediaType() (types.MediaType, error)
	Digest() (v1.Hash, error)
	Size() (int64, error)
}

type IndexAddendum struct {
	Add Appendable
	v1.Descriptor
}

func AppendLayers(base v1.Image, layers ...v1.Layer) (v1.Image, error) {
	panic("stub")
}

func Append(base v1.Image, adds ...Addendum) (v1.Image, error) {
	panic("stub")
}

func AppendManifests(base v1.ImageIndex, adds ...IndexAddendum) v1.ImageIndex {
	panic("stub")
}

func RemoveManifests(base v1.ImageIndex, matcher match.Matcher) v1.ImageIndex {
	panic("stub")
}

func Config(base v1.Image, cfg v1.Config) (v1.Image, error) {
	panic("stub")
}

func Subject(f partial.WithRawManifest, subject v1.Descriptor) partial.WithRawManifest {
	panic("stub")
}

func Annotations(f partial.WithRawManifest, anns map[string]string) partial.WithRawManifest {
	panic("stub")
}

func (a arbitraryRawManifest) RawManifest() ([]byte, error) {
	panic("stub")
}

func ConfigFile(base v1.Image, cfg *v1.ConfigFile) (v1.Image, error) {
	panic("stub")
}

func CreatedAt(base v1.Image, created v1.Time) (v1.Image, error) {
	panic("stub")
}

func Extract(img v1.Image) io.ReadCloser {
	panic("stub")
}

func Time(img v1.Image, t time.Time) (v1.Image, error) {
	panic("stub")
}

func Canonical(img v1.Image) (v1.Image, error) {
	panic("stub")
}

func MediaType(img v1.Image, mt types.MediaType) v1.Image {
	panic("stub")
}

func ConfigMediaType(img v1.Image, mt types.MediaType) v1.Image {
	panic("stub")
}

func IndexMediaType(idx v1.ImageIndex, mt types.MediaType) v1.ImageIndex {
	panic("stub")
}

func Rebase(orig, oldBase, newBase v1.Image) (v1.Image, error) {
	panic("stub")
}

type Embedme interface{}
