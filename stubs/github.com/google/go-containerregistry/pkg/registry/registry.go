package registry

import (
	"context"
	"io"
	"log"
	"net/http"
	"sync"

	"net/http/httptest"

	v1 "github.com/google/go-containerregistry/pkg/v1"
)

type blobHandler interface {
	Get(ctx context.Context, repo string, h v1.Hash) (io.ReadCloser, error)
}

type blobPutHandler interface {
	Put(ctx context.Context, repo string, h v1.Hash, rc io.ReadCloser) error
}

type memHandler struct {
	m    map[string][]byte
	lock sync.Mutex
}

type blobStatHandler interface {
	Stat(ctx context.Context, repo string, h v1.Hash) (int64, error)
}

type blobDeleteHandler interface {
	Delete(ctx context.Context, repo string, h v1.Hash) error
}

type redirectError struct {
	Location string
	Code     int
}

type blobs struct {
	blobHandler blobHandler
	uploads     map[string][]byte
	lock        sync.Mutex
	log         *log.Logger
}

func (e redirectError) Error() string {
	panic("stub")
}

func (m *memHandler) Stat(_ context.Context, _ string, h v1.Hash) (int64, error) {
	panic("stub")
}

func (m *memHandler) Get(_ context.Context, _ string, h v1.Hash) (io.ReadCloser, error) {
	panic("stub")
}

func (m *memHandler) Put(_ context.Context, _ string, h v1.Hash, rc io.ReadCloser) error {
	panic("stub")
}

func (m *memHandler) Delete(_ context.Context, _ string, h v1.Hash) error {
	panic("stub")
}

type regError struct {
	Status  int
	Code    string
	Message string
}

func (r *regError) Write(resp http.ResponseWriter) error {
	panic("stub")
}

type manifests struct {
	manifests map[string]map[string]manifest
	lock      sync.Mutex
	log       *log.Logger
}

type catalog struct{ Repos []string }

type listTags struct {
	Name string
	Tags []string
}

type manifest struct {
	contentType string
	blob        []byte
}

type registry struct {
	log              *log.Logger
	blobs            blobs
	manifests        manifests
	referrersEnabled bool
}

type Option func(r *registry)

func New(opts ...Option) http.Handler {
	panic("stub")
}

func Logger(l *log.Logger) Option {
	panic("stub")
}

func WithReferrersSupport(enabled bool) Option {
	panic("stub")
}

func TLS(domain string) (*httptest.Server, error) {
	panic("stub")
}

type Embedme interface{}
