package cache

import (
	"io"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type lazyLayer struct {
	inner v1.Layer
	c     Cache
}

type imageIndex struct {
	inner v1.ImageIndex
	c     Cache
}

type Cache interface {
	Put(v1.Layer) (v1.Layer, error)
	Get(v1.Hash) (v1.Layer, error)
	Delete(v1.Hash) error
}

type image struct {
	v1.Image
	c Cache
}

func Image(i v1.Image, c Cache) v1.Image {
	panic("stub")
}

func (i *image) Layers() ([]v1.Layer, error) {
	panic("stub")
}

func (l *lazyLayer) Compressed() (io.ReadCloser, error) {
	panic("stub")
}

func (l *lazyLayer) Uncompressed() (io.ReadCloser, error) {
	panic("stub")
}

func (l *lazyLayer) Size() (int64, error) {
	panic("stub")
}

func (l *lazyLayer) DiffID() (v1.Hash, error) {
	panic("stub")
}

func (l *lazyLayer) Digest() (v1.Hash, error) {
	panic("stub")
}

func (l *lazyLayer) MediaType() (types.MediaType, error) {
	panic("stub")
}

func (i *image) LayerByDigest(h v1.Hash) (v1.Layer, error) {
	panic("stub")
}

func (i *image) LayerByDiffID(h v1.Hash) (v1.Layer, error) {
	panic("stub")
}

func ImageIndex(ii v1.ImageIndex, c Cache) v1.ImageIndex {
	panic("stub")
}

func (ii *imageIndex) MediaType() (types.MediaType, error) {
	panic("stub")
}

func (ii *imageIndex) Digest() (v1.Hash, error) {
	panic("stub")
}

func (ii *imageIndex) Size() (int64, error) {
	panic("stub")
}

func (ii *imageIndex) IndexManifest() (*v1.IndexManifest, error) {
	panic("stub")
}

func (ii *imageIndex) RawManifest() ([]byte, error) {
	panic("stub")
}

func (ii *imageIndex) Image(h v1.Hash) (v1.Image, error) {
	panic("stub")
}

func (ii *imageIndex) ImageIndex(h v1.Hash) (v1.ImageIndex, error) {
	panic("stub")
}

type fscache struct{ path string }

type layer struct {
	v1.Layer
	path string
	digest
	diffID v1.Hash
}

type readcloser struct {
	t      io.Reader
	closes []func() error
}

func NewFilesystemCache(path string) Cache {
	panic("stub")
}

func (fs *fscache) Put(l v1.Layer) (v1.Layer, error) {
	panic("stub")
}

func (l *layer) Compressed() (io.ReadCloser, error) {
	panic("stub")
}

func (l *layer) Uncompressed() (io.ReadCloser, error) {
	panic("stub")
}

func (rc *readcloser) Read(b []byte) (int, error) {
	panic("stub")
}

func (rc *readcloser) Close() error {
	panic("stub")
}

func (fs *fscache) Get(h v1.Hash) (v1.Layer, error) {
	panic("stub")
}

func (fs *fscache) Delete(h v1.Hash) error {
	panic("stub")
}

type ro struct{ Cache }

func ReadOnly(c Cache) Cache {
	panic("stub")
}

type Embedme interface{}
