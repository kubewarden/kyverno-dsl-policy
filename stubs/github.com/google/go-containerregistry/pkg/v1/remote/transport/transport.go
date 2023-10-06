package transport

import (
	"context"
	"net/http"

	"github.com/google/go-containerregistry/internal/retry"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
)

type basicTransport struct {
	inner  http.RoundTripper
	auth   authn.Authenticator
	target string
}

func (bt *basicTransport) RoundTrip(in *http.Request) (*http.Response, error) {
	panic("stub")
}

type bearerTransport struct {
	inner    http.RoundTripper
	basic    authn.Authenticator
	bearer   authn.AuthConfig
	registry name.Registry
	realm    string
	service  string
	scopes   []string
	scheme   string
}

func (bt *bearerTransport) RoundTrip(in *http.Request) (*http.Response, error) {
	panic("stub")
}

const BlobUnknownErrorCode = "BLOB_UNKNOWN"

const BlobUploadInvalidErrorCode = "BLOB_UPLOAD_INVALID"

const BlobUploadUnknownErrorCode = "BLOB_UPLOAD_UNKNOWN"

const DigestInvalidErrorCode = "DIGEST_INVALID"

const ManifestBlobUnknownErrorCode = "MANIFEST_BLOB_UNKNOWN"

const ManifestInvalidErrorCode = "MANIFEST_INVALID"

const ManifestUnknownErrorCode = "MANIFEST_UNKNOWN"

const ManifestUnverifiedErrorCode = "MANIFEST_UNVERIFIED"

const NameInvalidErrorCode = "NAME_INVALID"

const NameUnknownErrorCode = "NAME_UNKNOWN"

const SizeInvalidErrorCode = "SIZE_INVALID"

const TagInvalidErrorCode = "TAG_INVALID"

const UnauthorizedErrorCode = "UNAUTHORIZED"

const DeniedErrorCode = "DENIED"

const UnsupportedErrorCode = "UNSUPPORTED"

const TooManyRequestsErrorCode = "TOOMANYREQUESTS"

const UnknownErrorCode = "UNKNOWN"

const UnavailableErrorCode = "UNAVAILABLE"

type Diagnostic struct {
	Code    ErrorCode
	Message string
	Detail  any
}

type Error struct {
	Errors     []Diagnostic
	StatusCode int
	Request    *http.Request
	rawBody    string
}

type ErrorCode string

func (e *Error) Error() string {
	panic("stub")
}

func (e *Error) Temporary() bool {
	panic("stub")
}

func (d Diagnostic) String() string {
	panic("stub")
}

func CheckError(resp *http.Response, codes ...int) error {
	panic("stub")
}

type logTransport struct{ inner http.RoundTripper }

func NewLogger(inner http.RoundTripper) http.RoundTripper {
	panic("stub")
}

func (t *logTransport) RoundTrip(in *http.Request) (out *http.Response, err error) {
	panic("stub")
}

const anonymous = "anonymous"

const basic = "basic"

const bearer = "bearer"

type multierrs []error

type challenge string

type pingResp struct {
	challenge  challenge
	parameters map[string]string
	scheme     string
}

func (c challenge) Canonical() challenge {
	panic("stub")
}

func (m multierrs) Error() string {
	panic("stub")
}

func (m multierrs) As(target any) bool {
	panic("stub")
}

func (m multierrs) Is(target error) bool {
	panic("stub")
}

type Backoff retry.Backoff

type retryTransport struct {
	inner     http.RoundTripper
	backoff   retry.Backoff
	predicate retry.Predicate
	codes     []int
}

type Option func(*options)

type options struct {
	backoff   retry.Backoff
	predicate retry.Predicate
	codes     []int
}

func WithRetryBackoff(backoff Backoff) Option {
	panic("stub")
}

func WithRetryPredicate(predicate func(error) bool) Option {
	panic("stub")
}

func WithRetryStatusCodes(codes ...int) Option {
	panic("stub")
}

func NewRetry(inner http.RoundTripper, opts ...Option) http.RoundTripper {
	panic("stub")
}

func (t *retryTransport) RoundTrip(in *http.Request) (out *http.Response, err error) {
	panic("stub")
}

type schemeTransport struct {
	scheme   string
	registry name.Registry
	inner    http.RoundTripper
}

func (st *schemeTransport) RoundTrip(in *http.Request) (*http.Response, error) {
	panic("stub")
}

const PullScope = "pull"

const PushScope = "push,pull"

const CatalogScope = "catalog"

type Wrapper struct{ inner http.RoundTripper }

func New(reg name.Registry, auth authn.Authenticator, t http.RoundTripper, scopes []string) (http.RoundTripper, error) {
	panic("stub")
}

func NewWithContext(ctx context.Context, reg name.Registry, auth authn.Authenticator, t http.RoundTripper, scopes []string) (http.RoundTripper, error) {
	panic("stub")
}

func (w *Wrapper) RoundTrip(in *http.Request) (*http.Response, error) {
	panic("stub")
}

var Version string

const defaultUserAgent = "go-containerregistry"

const moduleName = "github.com/google/go-containerregistry"

type userAgentTransport struct {
	inner http.RoundTripper
	ua    string
}

func NewUserAgent(inner http.RoundTripper, ua string) http.RoundTripper {
	panic("stub")
}

func (ut *userAgentTransport) RoundTrip(in *http.Request) (*http.Response, error) {
	panic("stub")
}

type Embedme interface{}
