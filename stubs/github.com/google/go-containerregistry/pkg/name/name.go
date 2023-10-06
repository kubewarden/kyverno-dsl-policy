package name

import _ "crypto/sha256"

const digestDelim = "@"

type Digest struct {
	Repository
	digest   string
	original string
}

func (d Digest) Context() Repository {
	panic("stub")
}

func (d Digest) Identifier() string {
	panic("stub")
}

func (d Digest) DigestStr() string {
	panic("stub")
}

func (d Digest) Name() string {
	panic("stub")
}

func (d Digest) String() string {
	panic("stub")
}

func NewDigest(name string, opts ...Option) (Digest, error) {
	panic("stub")
}

type ErrBadName struct{ info string }

func (e *ErrBadName) Error() string {
	panic("stub")
}

func (e *ErrBadName) Is(target error) bool {
	panic("stub")
}

func IsErrBadName(err error) bool {
	panic("stub")
}

const DefaultRegistry = "index.docker.io"

const defaultRegistryAlias = "docker.io"

const DefaultTag = "latest"

type OptionFn func() Option

type options struct {
	strict          bool
	insecure        bool
	defaultRegistry string
	defaultTag      string
}

type Option func(*options)

func StrictValidation(opts *options) {
	panic("stub")
}

func WeakValidation(opts *options) {
	panic("stub")
}

func Insecure(opts *options) {
	panic("stub")
}

func WithDefaultRegistry(r string) Option {
	panic("stub")
}

func WithDefaultTag(t string) Option {
	panic("stub")
}

type Reference interface {
	Context() Repository
	Identifier() string
	Name() string
	Scope(string) string
}

type stringConst string

func ParseReference(s string, opts ...Option) (Reference, error) {
	panic("stub")
}

func MustParseReference(s stringConst, opts ...Option) Reference {
	panic("stub")
}

type Registry struct {
	insecure bool
	registry string
}

func (r Registry) RegistryStr() string {
	panic("stub")
}

func (r Registry) Name() string {
	panic("stub")
}

func (r Registry) String() string {
	panic("stub")
}

func (r Registry) Scope(string) string {
	panic("stub")
}

func (r Registry) Scheme() string {
	panic("stub")
}

func NewRegistry(name string, opts ...Option) (Registry, error) {
	panic("stub")
}

func NewInsecureRegistry(name string, opts ...Option) (Registry, error) {
	panic("stub")
}

const defaultNamespace = "library"

const repositoryChars = "abcdefghijklmnopqrstuvwxyz0123456789_-./"

const regRepoDelimiter = "/"

type Repository struct {
	Registry
	repository string
}

func (r Repository) RepositoryStr() string {
	panic("stub")
}

func (r Repository) Name() string {
	panic("stub")
}

func (r Repository) String() string {
	panic("stub")
}

func (r Repository) Scope(action string) string {
	panic("stub")
}

func NewRepository(name string, opts ...Option) (Repository, error) {
	panic("stub")
}

func (r Repository) Tag(identifier string) Tag {
	panic("stub")
}

func (r Repository) Digest(identifier string) Digest {
	panic("stub")
}

const tagChars = "abcdefghijklmnopqrstuvwxyz0123456789_-.ABCDEFGHIJKLMNOPQRSTUVWXYZ"

const tagDelim = ":"

type Tag struct {
	Repository
	tag      string
	original string
}

func (t Tag) Context() Repository {
	panic("stub")
}

func (t Tag) Identifier() string {
	panic("stub")
}

func (t Tag) TagStr() string {
	panic("stub")
}

func (t Tag) Name() string {
	panic("stub")
}

func (t Tag) String() string {
	panic("stub")
}

func (t Tag) Scope(action string) string {
	panic("stub")
}

func NewTag(name string, opts ...Option) (Tag, error) {
	panic("stub")
}

type Embedme interface{}
