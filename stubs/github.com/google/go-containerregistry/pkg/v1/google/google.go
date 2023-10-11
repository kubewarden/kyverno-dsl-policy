package google

import (
	"context"
	"net/http"
	"os/exec"
	"sync"
	"time"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
)

const cloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

type gcloudSource struct{ exec func() *exec.Cmd }

type tokenSourceAuth struct{ Embedme }

type gcloudOutput struct{ Credential struct{} }

func NewEnvAuthenticator() (authn.Authenticator, error) {
	panic("stub")
}

func NewGcloudAuthenticator() (authn.Authenticator, error) {
	panic("stub")
}

func NewJSONKeyAuthenticator(serviceAccountJSON string) authn.Authenticator {
	panic("stub")
}

func NewTokenAuthenticator(serviceAccountJSON string, scope string) (authn.Authenticator, error) {
	panic("stub")
}

func NewTokenSourceAuthenticator(ts interface{}) authn.Authenticator {
	panic("stub")
}

func (tsa *tokenSourceAuth) Authorization() (*authn.AuthConfig, error) {
	panic("stub")
}

func (gs gcloudSource) Token() (interface{}, error) {
	panic("stub")
}

type googleKeychain struct {
	once sync.Once
	auth authn.Authenticator
}

func (gk *googleKeychain) Resolve(target authn.Resource) (authn.Authenticator, error) {
	panic("stub")
}

type ManifestInfo struct {
	Size      uint64
	MediaType string
	Created   time.Time
	Uploaded  time.Time
	Tags      []string
}

type Tags struct {
	Children  []string
	Manifests map[string]ManifestInfo
	Name      string
	Tags      []string
}

type Option func(*lister) error

type WalkFunc func(repo name.Repository, tags *Tags, err error) error

type lister struct {
	auth      authn.Authenticator
	transport http.RoundTripper
	repo      name.Repository
	client    *http.Client
	ctx       context.Context
	userAgent string
}

type rawManifestInfo struct {
	Size      string
	MediaType string
	Created   string
	Uploaded  string
	Tags      []string
}

func (m ManifestInfo) MarshalJSON() ([]byte, error) {
	panic("stub")
}

func (m *ManifestInfo) UnmarshalJSON(data []byte) error {
	panic("stub")
}

func List(repo name.Repository, options ...Option) (*Tags, error) {
	panic("stub")
}

func Walk(root name.Repository, walkFn WalkFunc, options ...Option) error {
	panic("stub")
}

func WithTransport(t http.RoundTripper) Option {
	panic("stub")
}

func WithAuth(auth authn.Authenticator) Option {
	panic("stub")
}

func WithAuthFromKeychain(keys authn.Keychain) Option {
	panic("stub")
}

func WithContext(ctx context.Context) Option {
	panic("stub")
}

func WithUserAgent(ua string) Option {
	panic("stub")
}

type Embedme interface{}
