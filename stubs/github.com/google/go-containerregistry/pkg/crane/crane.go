package crane

import (
	"context"
	"io"
	"net/http"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func Append(base v1.Image, paths ...string) (v1.Image, error) {
	panic("stub")
}

func Catalog(src string, opt ...Option) (res []string, err error) {
	panic("stub")
}

func Config(ref string, opt ...Option) ([]byte, error) {
	panic("stub")
}

func Copy(src, dst string, opt ...Option) error {
	panic("stub")
}

func Delete(src string, opt ...Option) error {
	panic("stub")
}

func Digest(ref string, opt ...Option) (string, error) {
	panic("stub")
}

func Export(img v1.Image, w io.Writer) error {
	panic("stub")
}

func Layer(filemap map[string][]byte) (v1.Layer, error) {
	panic("stub")
}

func Image(filemap map[string][]byte) (v1.Image, error) {
	panic("stub")
}

func Head(r string, opt ...Option) (*v1.Descriptor, error) {
	panic("stub")
}

func ListTags(src string, opt ...Option) ([]string, error) {
	panic("stub")
}

func Manifest(ref string, opt ...Option) ([]byte, error) {
	panic("stub")
}

type stringSet map[string]struct{}

func Optimize(src, dst string, prioritize []string, opt ...Option) error {
	panic("stub")
}

func (s stringSet) List() []string {
	panic("stub")
}

func (s stringSet) Intersection(rhs stringSet) stringSet {
	panic("stub")
}

type Option func(*Options)

type Options struct {
	Name      []name.Option
	Remote    []remote.Option
	Platform  *v1.Platform
	Keychain  authn.Keychain
	transport http.RoundTripper
	insecure  bool
}

func GetOptions(opts ...Option) Options {
	panic("stub")
}

func WithTransport(t http.RoundTripper) Option {
	panic("stub")
}

func Insecure(o *Options) {
	panic("stub")
}

func WithPlatform(platform *v1.Platform) Option {
	panic("stub")
}

func WithAuthFromKeychain(keys authn.Keychain) Option {
	panic("stub")
}

func WithAuth(auth authn.Authenticator) Option {
	panic("stub")
}

func WithUserAgent(ua string) Option {
	panic("stub")
}

func WithNondistributable() Option {
	panic("stub")
}

func WithContext(ctx context.Context) Option {
	panic("stub")
}

const iWasADigestTag = "i-was-a-digest"

func Pull(src string, opt ...Option) (v1.Image, error) {
	panic("stub")
}

func Save(img v1.Image, src, path string) error {
	panic("stub")
}

func MultiSave(imgMap map[string]v1.Image, path string, opt ...Option) error {
	panic("stub")
}

func PullLayer(ref string, opt ...Option) (v1.Layer, error) {
	panic("stub")
}

func SaveLegacy(img v1.Image, src, path string) error {
	panic("stub")
}

func MultiSaveLegacy(imgMap map[string]v1.Image, path string) error {
	panic("stub")
}

func SaveOCI(img v1.Image, path string) error {
	panic("stub")
}

func MultiSaveOCI(imgMap map[string]v1.Image, path string) error {
	panic("stub")
}

func Load(path string, opt ...Option) (v1.Image, error) {
	panic("stub")
}

func LoadTag(path, tag string, opt ...Option) (v1.Image, error) {
	panic("stub")
}

func Push(img v1.Image, dst string, opt ...Option) error {
	panic("stub")
}

func Upload(layer v1.Layer, repo string, opt ...Option) error {
	panic("stub")
}

func Tag(img, tag string, opt ...Option) error {
	panic("stub")
}

type Embedme interface{}
