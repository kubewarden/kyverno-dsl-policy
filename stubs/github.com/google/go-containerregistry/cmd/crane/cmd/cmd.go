package cmd

import (
	"github.com/google/go-containerregistry/pkg/crane"

	"net/http"

	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/partial"
)

func NewCmdAppend(options *[]crane.Option) interface{} {
	panic("stub")
}

type credentials struct {
	Username string
	Secret   string
}

type loginOptions struct {
	serverAddress string
	user          string
	password      string
	passwordStdin bool
}

func NewCmdAuth(options []crane.Option, argv ...string) interface{} {
	panic("stub")
}

func NewCmdAuthGet(options []crane.Option, argv ...string) interface{} {
	panic("stub")
}

func NewCmdAuthLogin(argv ...string) interface{} {
	panic("stub")
}

func NewCmdBlob(options *[]crane.Option) interface{} {
	panic("stub")
}

func NewCmdCatalog(options *[]crane.Option, argv ...string) interface{} {
	panic("stub")
}

func NewCmdConfig(options *[]crane.Option) interface{} {
	panic("stub")
}

func NewCmdCopy(options *[]crane.Option) interface{} {
	panic("stub")
}

func NewCmdDelete(options *[]crane.Option) interface{} {
	panic("stub")
}

func NewCmdDigest(options *[]crane.Option) interface{} {
	panic("stub")
}

func NewCmdExport(options *[]crane.Option) interface{} {
	panic("stub")
}

type remoteIndex interface {
	Manifests() ([]partial.Describable, error)
}

func NewCmdFlatten(options *[]crane.Option) interface{} {
	panic("stub")
}

func NewCmdIndex(options *[]crane.Option) interface{} {
	panic("stub")
}

func NewCmdIndexFilter(options *[]crane.Option) interface{} {
	panic("stub")
}

func NewCmdIndexAppend(options *[]crane.Option) interface{} {
	panic("stub")
}

func NewCmdList(options *[]crane.Option) interface{} {
	panic("stub")
}

func NewCmdManifest(options *[]crane.Option) interface{} {
	panic("stub")
}

func NewCmdMutate(options *[]crane.Option) interface{} {
	panic("stub")
}

func NewCmdOptimize(options *[]crane.Option) interface{} {
	panic("stub")
}

func NewCmdPull(options *[]crane.Option) interface{} {
	panic("stub")
}

func NewCmdPush(options *[]crane.Option) interface{} {
	panic("stub")
}

func NewCmdRebase(options *[]crane.Option) interface{} {
	panic("stub")
}

const use = "crane"

const short = "Crane is a tool for managing container images"

type headerTransport struct {
	httpHeaders map[string]string
	inner       http.RoundTripper
}

func New(use, short string, options []crane.Option) interface{} {
	panic("stub")
}

func (ht *headerTransport) RoundTrip(in *http.Request) (*http.Response, error) {
	panic("stub")
}

func NewCmdTag(options *[]crane.Option) interface{} {
	panic("stub")
}

type platformsValue struct{ platforms []v1.Platform }

type platformValue struct{ platform *v1.Platform }

func (ps *platformsValue) Set(platform string) error {
	panic("stub")
}

func (ps *platformsValue) String() string {
	panic("stub")
}

func (ps *platformsValue) Type() string {
	panic("stub")
}

func (pv *platformValue) Set(platform string) error {
	panic("stub")
}

func (pv *platformValue) String() string {
	panic("stub")
}

func (pv *platformValue) Type() string {
	panic("stub")
}

func NewCmdValidate(options *[]crane.Option) interface{} {
	panic("stub")
}

var Version string

func NewCmdVersion() interface{} {
	panic("stub")
}

type Embedme interface{}
