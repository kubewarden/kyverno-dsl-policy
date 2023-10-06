package v1

import (
	"hash"
	"io"
	"time"

	"github.com/google/go-containerregistry/pkg/v1/types"
)

type ConfigFile struct {
	Architecture  string
	Author        string
	Container     string
	Created       Time
	DockerVersion string
	History       []History
	OS            string
	RootFS        RootFS
	Config        Config
	OSVersion     string
	Variant       string
	OSFeatures    []string
}

type History struct {
	Author     string
	Created    Time
	CreatedBy  string
	Comment    string
	EmptyLayer bool
}

type Time struct{ time.Time }

type RootFS struct {
	Type    string
	DiffIDs []Hash
}

type HealthConfig struct {
	Test        []string
	Interval    time.Duration
	Timeout     time.Duration
	StartPeriod time.Duration
	Retries     int
}

type Config struct {
	AttachStderr    bool
	AttachStdin     bool
	AttachStdout    bool
	Cmd             []string
	Healthcheck     *HealthConfig
	Domainname      string
	Entrypoint      []string
	Env             []string
	Hostname        string
	Image           string
	Labels          map[string]string
	OnBuild         []string
	OpenStdin       bool
	StdinOnce       bool
	Tty             bool
	User            string
	Volumes         map[string]struct{}
	WorkingDir      string
	ExposedPorts    map[string]struct{}
	ArgsEscaped     bool
	NetworkDisabled bool
	MacAddress      string
	StopSignal      string
	Shell           []string
}

func (cf *ConfigFile) Platform() *Platform {
	panic("stub")
}

func (t *Time) DeepCopyInto(out *Time) {
	panic("stub")
}

func ParseConfigFile(r io.Reader) (*ConfigFile, error) {
	panic("stub")
}

type Hash struct {
	Algorithm string
	Hex       string
}

func (h Hash) String() string {
	panic("stub")
}

func NewHash(s string) (Hash, error) {
	panic("stub")
}

func (h Hash) MarshalJSON() ([]byte, error) {
	panic("stub")
}

func (h *Hash) UnmarshalJSON(data []byte) error {
	panic("stub")
}

func (h Hash) MarshalText() (text []byte, err error) {
	panic("stub")
}

func (h *Hash) UnmarshalText(text []byte) error {
	panic("stub")
}

func Hasher(name string) (hash.Hash, error) {
	panic("stub")
}

func SHA256(r io.Reader) (Hash, int64, error) {
	panic("stub")
}

type Image interface {
	Layers() ([]Layer, error)
	MediaType() (types.MediaType, error)
	Size() (int64, error)
	ConfigName() (Hash, error)
	ConfigFile() (*ConfigFile, error)
	RawConfigFile() ([]byte, error)
	Digest() (Hash, error)
	Manifest() (*Manifest, error)
	RawManifest() ([]byte, error)
	LayerByDigest(Hash) (Layer, error)
	LayerByDiffID(Hash) (Layer, error)
}

type ImageIndex interface {
	MediaType() (types.MediaType, error)
	Digest() (Hash, error)
	Size() (int64, error)
	IndexManifest() (*IndexManifest, error)
	RawManifest() ([]byte, error)
	Image(Hash) (Image, error)
	ImageIndex(Hash) (ImageIndex, error)
}

type Layer interface {
	Digest() (Hash, error)
	DiffID() (Hash, error)
	Compressed() (io.ReadCloser, error)
	Uncompressed() (io.ReadCloser, error)
	Size() (int64, error)
	MediaType() (types.MediaType, error)
}

type Manifest struct {
	SchemaVersion int64
	MediaType     types.MediaType
	Config        Descriptor
	Layers        []Descriptor
	Annotations   map[string]string
	Subject       *Descriptor
}

type IndexManifest struct {
	SchemaVersion int64
	MediaType     types.MediaType
	Manifests     []Descriptor
	Annotations   map[string]string
	Subject       *Descriptor
}

type Descriptor struct {
	MediaType    types.MediaType
	Size         int64
	Digest       Hash
	Data         []byte
	URLs         []string
	Annotations  map[string]string
	Platform     *Platform
	ArtifactType string
}

func ParseManifest(r io.Reader) (*Manifest, error) {
	panic("stub")
}

func ParseIndexManifest(r io.Reader) (*IndexManifest, error) {
	panic("stub")
}

type Platform struct {
	Architecture string
	OS           string
	OSVersion    string
	OSFeatures   []string
	Variant      string
	Features     []string
}

func (p Platform) String() string {
	panic("stub")
}

func ParsePlatform(s string) (*Platform, error) {
	panic("stub")
}

func (p Platform) Equals(o Platform) bool {
	panic("stub")
}

func (p Platform) Satisfies(spec Platform) bool {
	panic("stub")
}

type Update struct {
	Total    int64
	Complete int64
	Error    error
}

type Embedme interface{}
