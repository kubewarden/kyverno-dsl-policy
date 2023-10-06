package gcrane

import (
	"context"
	"net/http"

	"github.com/google/go-containerregistry/internal/retry"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/google"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

type task struct {
	digest   string
	manifest google.ManifestInfo
	oldRepo  name.Repository
	newRepo  name.Repository
}

type copier struct {
	srcRepo name.Repository
	dstRepo name.Repository
	tasks   chan task
	opt     *options
}

func GCRBackoff() retry.Backoff {
	panic("stub")
}

func Copy(src, dst string, opts ...Option) error {
	panic("stub")
}

func CopyRepository(ctx context.Context, src, dst string, opts ...Option) error {
	panic("stub")
}

type Option func(*options)

type options struct {
	jobs   int
	remote []remote.Option
	google []google.Option
	crane  []crane.Option
}

func WithJobs(jobs int) Option {
	panic("stub")
}

func WithTransport(t http.RoundTripper) Option {
	panic("stub")
}

func WithUserAgent(ua string) Option {
	panic("stub")
}

func WithContext(ctx context.Context) Option {
	panic("stub")
}

func WithKeychain(keys authn.Keychain) Option {
	panic("stub")
}

func WithAuth(auth authn.Authenticator) Option {
	panic("stub")
}

type Embedme interface{}
