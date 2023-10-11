package registryclient

import (
	"context"
	"net/http"

	"github.com/google/go-containerregistry/pkg/authn"

	gcrremote "github.com/google/go-containerregistry/pkg/v1/remote"
)

type autoRefreshSecrets struct {
	lister           interface{}
	imagePullSecrets []string
}

func NewAutoRefreshSecretsKeychain(lister interface{}, imagePullSecrets ...string) (authn.Keychain, error) {
	panic("stub")
}

func (kc *autoRefreshSecrets) Resolve(resource authn.Resource) (authn.Authenticator, error) {
	panic("stub")
}

type Client interface {
	Keychain() authn.Keychain
	getTransport() http.RoundTripper
	FetchImageDescriptor(context.Context, string) (*gcrremote.Descriptor, error)
	BuildRemoteOption(context.Context) interface{}
}

type config struct {
	keychain  []authn.Keychain
	transport *http.Transport
	tracing   bool
}

type Option func(*config) error

type client struct {
	keychain  authn.Keychain
	transport http.RoundTripper
}

func New(options ...Option) (Client, error) {
	panic("stub")
}

func NewOrDie(options ...Option) Client {
	panic("stub")
}

func WithKeychainPullSecrets(lister interface{}, imagePullSecrets ...string) Option {
	panic("stub")
}

func WithCredentialHelpers(credentialHelpers ...string) Option {
	panic("stub")
}

func WithAllowInsecureRegistry() Option {
	panic("stub")
}

func WithLocalKeychain() Option {
	panic("stub")
}

func WithTracing() Option {
	panic("stub")
}

func (c *client) BuildRemoteOption(ctx context.Context) interface{} {
	panic("stub")
}

func (c *client) FetchImageDescriptor(ctx context.Context, imageRef string) (*gcrremote.Descriptor, error) {
	panic("stub")
}

func (c *client) Keychain() authn.Keychain {
	panic("stub")
}

type Embedme interface{}
