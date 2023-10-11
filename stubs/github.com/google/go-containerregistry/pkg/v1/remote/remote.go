package remote

import (
	"context"
	"io"
	"net/http"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"

	"sync"

	"github.com/google/go-containerregistry/internal/retry"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/partial"
	"github.com/google/go-containerregistry/pkg/v1/types"
)

type catalog struct{ Repos []string }

func CatalogPage(target name.Registry, last string, n int, options ...Option) ([]string, error) {
	panic("stub")
}

func Catalog(ctx context.Context, target name.Registry, options ...Option) ([]string, error) {
	panic("stub")
}

func CheckPushPermission(ref name.Reference, kc authn.Keychain, t http.RoundTripper) error {
	panic("stub")
}

func Delete(ref name.Reference, options ...Option) error {
	panic("stub")
}

type ErrSchema1 struct{ schema string }

type Descriptor struct {
	fetcher
	v1.Descriptor
	Manifest []byte
	platform v1.Platform
}

type fetcher struct {
	Ref     name.Reference
	Client  *http.Client
	context context.Context
}

func (e *ErrSchema1) Error() string {
	panic("stub")
}

func (d *Descriptor) RawManifest() ([]byte, error) {
	panic("stub")
}

func Get(ref name.Reference, options ...Option) (*Descriptor, error) {
	panic("stub")
}

func Head(ref name.Reference, options ...Option) (*v1.Descriptor, error) {
	panic("stub")
}

func (d *Descriptor) Image() (v1.Image, error) {
	panic("stub")
}

func (d *Descriptor) ImageIndex() (v1.ImageIndex, error) {
	panic("stub")
}

type remoteImage struct {
	fetcher
	manifestLock sync.Mutex
	manifest     []byte
	configLock   sync.Mutex
	config       []byte
	mediaType    types.MediaType
	descriptor   *v1.Descriptor
}

type remoteImageLayer struct {
	ri     *remoteImage
	digest v1.Hash
}

func (r *remoteImage) ArtifactType() (string, error) {
	panic("stub")
}

func Image(ref name.Reference, options ...Option) (v1.Image, error) {
	panic("stub")
}

func (r *remoteImage) MediaType() (types.MediaType, error) {
	panic("stub")
}

func (r *remoteImage) RawManifest() ([]byte, error) {
	panic("stub")
}

func (r *remoteImage) RawConfigFile() ([]byte, error) {
	panic("stub")
}

func (r *remoteImage) Descriptor() (*v1.Descriptor, error) {
	panic("stub")
}

func (rl *remoteImageLayer) Digest() (v1.Hash, error) {
	panic("stub")
}

func (rl *remoteImageLayer) Compressed() (io.ReadCloser, error) {
	panic("stub")
}

func (rl *remoteImageLayer) Manifest() (*v1.Manifest, error) {
	panic("stub")
}

func (rl *remoteImageLayer) MediaType() (types.MediaType, error) {
	panic("stub")
}

func (rl *remoteImageLayer) Size() (int64, error) {
	panic("stub")
}

func (rl *remoteImageLayer) ConfigFile() (*v1.ConfigFile, error) {
	panic("stub")
}

func (rl *remoteImageLayer) DiffID() (v1.Hash, error) {
	panic("stub")
}

func (rl *remoteImageLayer) Descriptor() (*v1.Descriptor, error) {
	panic("stub")
}

func (rl *remoteImageLayer) Exists() (bool, error) {
	panic("stub")
}

func (r *remoteImage) LayerByDigest(h v1.Hash) (partial.CompressedLayer, error) {
	panic("stub")
}

type remoteIndex struct {
	fetcher
	manifestLock sync.Mutex
	manifest     []byte
	mediaType    types.MediaType
	descriptor   *v1.Descriptor
}

func Index(ref name.Reference, options ...Option) (v1.ImageIndex, error) {
	panic("stub")
}

func (r *remoteIndex) MediaType() (types.MediaType, error) {
	panic("stub")
}

func (r *remoteIndex) Digest() (v1.Hash, error) {
	panic("stub")
}

func (r *remoteIndex) Size() (int64, error) {
	panic("stub")
}

func (r *remoteIndex) RawManifest() ([]byte, error) {
	panic("stub")
}

func (r *remoteIndex) IndexManifest() (*v1.IndexManifest, error) {
	panic("stub")
}

func (r *remoteIndex) Image(h v1.Hash) (v1.Image, error) {
	panic("stub")
}

func (r *remoteIndex) Descriptor() (*v1.Descriptor, error) {
	panic("stub")
}

func (r *remoteIndex) ImageIndex(h v1.Hash) (v1.ImageIndex, error) {
	panic("stub")
}

func (r *remoteIndex) Layer(h v1.Hash) (v1.Layer, error) {
	panic("stub")
}

func (r *remoteIndex) Manifests() ([]partial.Describable, error) {
	panic("stub")
}

type remoteLayer struct {
	fetcher
	digest v1.Hash
}

func (rl *remoteLayer) Compressed() (io.ReadCloser, error) {
	panic("stub")
}

func (rl *remoteLayer) Size() (int64, error) {
	panic("stub")
}

func (rl *remoteLayer) Digest() (v1.Hash, error) {
	panic("stub")
}

func (rl *remoteLayer) MediaType() (types.MediaType, error) {
	panic("stub")
}

func (rl *remoteLayer) Exists() (bool, error) {
	panic("stub")
}

func Layer(ref name.Digest, options ...Option) (v1.Layer, error) {
	panic("stub")
}

type tags struct {
	Name string
	Tags []string
}

func ListWithContext(ctx context.Context, repo name.Repository, options ...Option) ([]string, error) {
	panic("stub")
}

func List(repo name.Repository, options ...Option) ([]string, error) {
	panic("stub")
}

type MountableLayer struct {
	v1.Layer
	Reference name.Reference
}

type mountableImage struct {
	v1.Image
	Reference name.Reference
}

func (ml *MountableLayer) Descriptor() (*v1.Descriptor, error) {
	panic("stub")
}

func (ml *MountableLayer) Exists() (bool, error) {
	panic("stub")
}

func (mi *mountableImage) Layers() ([]v1.Layer, error) {
	panic("stub")
}

func (mi *mountableImage) LayerByDigest(d v1.Hash) (v1.Layer, error) {
	panic("stub")
}

func (mi *mountableImage) LayerByDiffID(d v1.Hash) (v1.Layer, error) {
	panic("stub")
}

func (mi *mountableImage) Descriptor() (*v1.Descriptor, error) {
	panic("stub")
}

func (mi *mountableImage) ConfigLayer() (v1.Layer, error) {
	panic("stub")
}

func MultiWrite(m map[name.Reference]Taggable, options ...Option) (rerr error) {
	panic("stub")
}

const defaultJobs = 4

const defaultPageSize = 1000

type Option func(*options) error

type options struct {
	auth                           authn.Authenticator
	keychain                       authn.Keychain
	transport                      http.RoundTripper
	platform                       v1.Platform
	context                        context.Context
	jobs                           int
	userAgent                      string
	allowNondistributableArtifacts bool
	updates                        chan<- v1.Update
	pageSize                       int
	retryBackoff                   Backoff
	retryPredicate                 retry.Predicate
	filter                         map[string]string
}

type Backoff retry.Backoff

func WithTransport(t http.RoundTripper) Option {
	panic("stub")
}

func WithAuth(auth authn.Authenticator) Option {
	panic("stub")
}

func WithAuthFromKeychain(keys authn.Keychain) Option {
	panic("stub")
}

func WithPlatform(p v1.Platform) Option {
	panic("stub")
}

func WithContext(ctx context.Context) Option {
	panic("stub")
}

func WithJobs(jobs int) Option {
	panic("stub")
}

func WithUserAgent(ua string) Option {
	panic("stub")
}

func WithNondistributable(o *options) error {
	panic("stub")
}

func WithProgress(updates chan<- v1.Update) Option {
	panic("stub")
}

func WithPageSize(size int) Option {
	panic("stub")
}

func WithRetryBackoff(backoff Backoff) Option {
	panic("stub")
}

func WithRetryPredicate(predicate retry.Predicate) Option {
	panic("stub")
}

func WithFilter(key string, value string) Option {
	panic("stub")
}

type progress struct {
	sync.Mutex
	updates    chan<- v1.Update
	lastUpdate *v1.Update
}

type progressReader struct {
	rc       io.ReadCloser
	count    *int64
	progress *progress
}

func (r *progressReader) Read(b []byte) (int, error) {
	panic("stub")
}

func (r *progressReader) Close() error {
	panic("stub")
}

func Referrers(d name.Digest, options ...Option) (*v1.IndexManifest, error) {
	panic("stub")
}

type Taggable interface {
	RawManifest() ([]byte, error)
}

type writer struct {
	repo      name.Repository
	client    *http.Client
	progress  *progress
	backoff   Backoff
	predicate retry.Predicate
}

type withLayer interface {
	Layer(v1.Hash) (v1.Layer, error)
}

type fallbackTaggable struct{ im v1.IndexManifest }

type withMediaType interface {
	MediaType() (types.MediaType, error)
}

func Write(ref name.Reference, img v1.Image, options ...Option) (rerr error) {
	panic("stub")
}

func (f fallbackTaggable) RawManifest() ([]byte, error) {
	panic("stub")
}

func (f fallbackTaggable) MediaType() (types.MediaType, error) {
	panic("stub")
}

func WriteIndex(ref name.Reference, ii v1.ImageIndex, options ...Option) (rerr error) {
	panic("stub")
}

func WriteLayer(repo name.Repository, layer v1.Layer, options ...Option) (rerr error) {
	panic("stub")
}

func Tag(tag name.Tag, t Taggable, options ...Option) error {
	panic("stub")
}

func Put(ref name.Reference, t Taggable, options ...Option) error {
	panic("stub")
}

type Embedme interface{}
