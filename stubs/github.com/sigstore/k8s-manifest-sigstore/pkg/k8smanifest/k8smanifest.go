package k8smanifest

import (
	"crypto/x509"

	_ "embed"

	kubeutil "github.com/sigstore/k8s-manifest-sigstore/pkg/util/kubeutil"
	"github.com/sigstore/k8s-manifest-sigstore/pkg/util/mapnode"

	cryptox509 "crypto/x509"
	"time"
)

type SignatureNotFoundError struct{ *K8sManifestError }

type SignatureVerificationError struct{ *K8sManifestError }

type K8sManifestError struct {
	error
	defaultMessage string
}

type MessageNotFoundError struct{ *K8sManifestError }

func (e *K8sManifestError) Error() string {
	panic("stub")
}

func NewSignatureNotFoundError(err error) *SignatureNotFoundError {
	panic("stub")
}

func NewMessageNotFoundError(err error) *MessageNotFoundError {
	panic("stub")
}

func NewSignatureVerificationError(err error) *SignatureVerificationError {
	panic("stub")
}

func IsSignatureNotFoundError(err error) bool {
	panic("stub")
}

func IsMessageNotFoundError(err error) bool {
	panic("stub")
}

func IsSignatureVerificationError(err error) bool {
	panic("stub")
}

var defaultConfigBytes []byte

type cosignVerifyOption struct {
	Certificate      string
	CertificateChain string
	RekorURL         string
	OIDCIssuer       string
	RootCerts        *x509.CertPool
	AllowInsecure    bool
}

type AnnotationConfig struct {
	AnnotationKeyDomain       string
	ResourceBundleRefBaseName string
	SignatureBaseName         string
	CertificateBaseName       string
	MessageBaseName           string
	BundleBaseName            string
}

type cosignSignOption struct {
	RekorURL      string
	NoTlogUpload  bool
	AllowInsecure bool
	Force         bool
}

type VerifyResourceOption struct {
	commonOption
	verifyOption
	cosignVerifyOption
	SkipObjects           ObjectReferenceList
	Provenance            bool
	DisableDryRun         bool
	CheckDryRunForApply   bool
	CheckMutatingResource bool
	DryRunNamespace       string
}

type verifyOption struct {
	IgnoreFields                ObjectFieldBindingList
	Signers                     SignerList
	MaxResourceManifestNum      int
	KeyPath                     string
	ResourceBundleRef           string
	SignatureResourceRef        string
	ProvenanceResourceRef       string
	UseCache                    bool
	CacheDir                    string
	annotationKeyToIgnoreFields bool
}

type ObjectReferenceList []ObjectReference

type ObjectFieldBindingList []ObjectFieldBinding

type SignOption struct {
	commonOption
	cosignSignOption
	KeyPath           string
	ResourceBundleRef string
	CertPath          string
	Output            string
	UpdateAnnotation  bool
	ImageAnnotations  map[string]interface{}
	PassFunc          interface{}
	ApplySigConfigMap bool
	Tarball           *bool
	AppendSignature   bool
}

type ObjectFieldBinding struct {
	Fields  []string
	Objects ObjectReferenceList
}

type SignerList []string

type ObjectUserBinding struct {
	Users   []string
	Objects ObjectReferenceList
}

type commonOption struct{ AnnotationConfig }

type ObjectReference struct {
	Group     string
	Version   string
	Kind      string
	Name      string
	Namespace string
}

type VerifyManifestOption struct {
	commonOption
	verifyOption
	cosignVerifyOption
}

func (o *VerifyResourceOption) SetAnnotationIgnoreFields() {
	panic("stub")
}

func (o *VerifyManifestOption) SetAnnotationIgnoreFields() {
	panic("stub")
}

func (c AnnotationConfig) MessageAnnotationKey() string {
	panic("stub")
}

func (c AnnotationConfig) SignatureAnnotationKey(i int) string {
	panic("stub")
}

func (c AnnotationConfig) CertificateAnnotationKey(i int) string {
	panic("stub")
}

func (c AnnotationConfig) BundleAnnotationKey(i int) string {
	panic("stub")
}

func (c AnnotationConfig) ResourceBundleRefAnnotationKey() string {
	panic("stub")
}

func (c AnnotationConfig) AnnotationKeyMap(i int) map[string]string {
	panic("stub")
}

func (c AnnotationConfig) GetAllSignatureSets(annotations map[string]string) []map[string]string {
	panic("stub")
}

func (c AnnotationConfig) AnnotationKeyMask() []string {
	panic("stub")
}

func (c AnnotationConfig) AnnotationKeyIgnoreField() ObjectFieldBindingList {
	panic("stub")
}

func ObjectToReference(obj interface{}) ObjectReference {
	panic("stub")
}

func (l ObjectReferenceList) Match(obj interface{}) bool {
	panic("stub")
}

func (r ObjectReference) Match(obj interface{}) bool {
	panic("stub")
}

func (r ObjectReference) Equal(r2 ObjectReference) bool {
	panic("stub")
}

func (l ObjectFieldBindingList) Match(obj interface{}) (bool, []string) {
	panic("stub")
}

func (f ObjectFieldBinding) Match(obj interface{}) (bool, []string) {
	panic("stub")
}

func (l SignerList) Match(signerName string) bool {
	panic("stub")
}

func LoadVerifyManifestConfig(fpath string) (*VerifyManifestOption, error) {
	panic("stub")
}

func LoadVerifyResourceConfig(fpath string) (*VerifyResourceOption, error) {
	panic("stub")
}

func LoadVerifyResourceConfigFromResource(configPath, configField string) (*VerifyResourceOption, error) {
	panic("stub")
}

func GetConfigResource(configPath string) (interface{}, error) {
	panic("stub")
}

func GetMatchConditionFromConfigResource(configPath, matchField, inScopeObjectField string) (interface{}, *ObjectReferenceList, error) {
	panic("stub")
}

func (vo *VerifyResourceOption) AddDefaultConfig(defaultConfig *VerifyResourceOption) *VerifyResourceOption {
	panic("stub")
}

func LoadDefaultConfig() *VerifyResourceOption {
	panic("stub")
}

func AddDefaultConfig(vo *VerifyResourceOption) *VerifyResourceOption {
	panic("stub")
}

const ArtifactUnknown = ""

const ArtifactManifestImage = "manifestImage"

const ArtifactManifestResource = "manifestResource"

const ArtifactContainerImage = "containerImage"

const AttestationDataKeyName = "attestation"

const SBOMDataKeyName = "sbom"

type ProvenanceGetter interface {
	Get() ([]*Provenance, error)
}

type ResourceProvenanceGetter struct{ resourceRefString string }

type NotImplementedProvenanceGetter struct{}

type resourceName struct {
	PodName       string
	ContainerName string
}

type ArtifactType string

type ImageProvenanceGetter struct {
	resBundleRef  string
	imageHash     string
	cacheEnabled  bool
	allowInsecure bool
}

type DigestSet map[string]string

type RecursiveImageProvenanceGetter struct {
	object                        interface{}
	manifestResourceBundleRef     string
	manifestProvenanceResourceRef string
	cacheEnabled                  bool
	allowInsecure                 bool
}

type ProvenanceMaterial struct {
	URI    string
	Digest DigestSet
}

type Provenance struct {
	ResourceName         *resourceName
	RawAttestation       string
	RawSBOM              string
	Artifact             string
	ArtifactType         ArtifactType
	Hash                 string
	AttestationLogIndex  *int
	AttestationMaterials []ProvenanceMaterial
	SBOMRef              string
	ConfigMapRef         string
}

type rekorCLIGetCmdOutput struct {
	Attestation     string
	AttestationType string
	Body            interface{}
	LogIndex        int
	IntegratedTime  int64
	UUID            string
	LogID           string
}

func NewProvenanceGetter(obj interface{}, sigRef, imageHash, provResRef string, allowInsecure bool) ProvenanceGetter {
	panic("stub")
}

func (g *RecursiveImageProvenanceGetter) Get() ([]*Provenance, error) {
	panic("stub")
}

func (g *ImageProvenanceGetter) Get() ([]*Provenance, error) {
	panic("stub")
}

func (g *ResourceProvenanceGetter) Get() ([]*Provenance, error) {
	panic("stub")
}

func (g *NotImplementedProvenanceGetter) Get() ([]*Provenance, error) {
	panic("stub")
}

func ParseAttestation(attestationStr string) (interface{}, interface{}, []ProvenanceMaterial, error) {
	panic("stub")
}

func GenerateIntotoAttestationCurlCommand(logIndex int) string {
	panic("stub")
}

func GenerateIntotoAttestationKubectlCommand(resourceRef string) string {
	panic("stub")
}

func GenerateSBOMDownloadCommand(resBundleRef string) string {
	panic("stub")
}

func GenerateSBOMKubectlCommand(resourceRef string) string {
	panic("stub")
}

const DefaultAnnotationKeyDomain = "cosign.sigstore.dev"

const defaultMessageAnnotationBaseName = "message"

const defaultSignatureAnnotationBaseName = "signature"

const defaultCertificateAnnotationBaseName = "certificate"

const defaultBundleAnnotationBaseName = "bundle"

const defaultResourceBundleRefAnnotationBaseName = "resourceBundleRef"

type CosignSignConfig struct {
	RekorURL      string
	NoTlogUpload  bool
	AllowInsecure bool
	Force         bool
}

type Signer interface {
	Sign(inputDir, output string, imageAnnotations map[string]interface{}) ([]byte, error)
}

type ImageSigner struct {
	AnnotationConfig AnnotationConfig
	tarball          bool
	resBundleRef     string
	prikeyPath       *string
	certPath         *string
	passFunc         interface{}
	CosignSignConfig
}

type BlobSigner struct {
	AnnotationConfig   AnnotationConfig
	createSigConfigMap bool
	appendSig          bool
	doApply            bool
	tarball            bool
	prikeyPath         *string
	certPath           *string
	passFunc           interface{}
	CosignSignConfig
}

func Sign(inputDir string, so *SignOption) ([]byte, error) {
	panic("stub")
}

func NewSigner(resBundleRef, keyPath, certPath, output string, appendSig, doApply, tarball bool, cosignSignConfig CosignSignConfig, AnnotationConfig AnnotationConfig, pf interface{}) Signer {
	panic("stub")
}

func (s *ImageSigner) Sign(inputDir, output string, imageAnnotations map[string]interface{}) ([]byte, error) {
	panic("stub")
}

func (s *BlobSigner) Sign(inputDir, output string, imageAnnotations map[string]interface{}) ([]byte, error) {
	panic("stub")
}

func K8sResourceRef2FileName(resRef string) string {
	panic("stub")
}

const SigRefEmbeddedInAnnotation = "__embedded_in_annotation__"

type ImageManifestFetcher struct {
	resBundleRefString     string
	AnnotationConfig       AnnotationConfig
	ignoreFields           []string
	maxResourceManifestNum int
	cacheEnabled           bool
	allowInsecure          bool
}

type ManifestFetcher interface {
	Fetch(objYAMLBytes []byte) ([][]byte, string, error)
}

type SignatureVerifier interface {
	Verify() (bool, string, *int64, error)
}

type verificationIdentity struct {
	path string
	name string
}

type VerifyResult struct {
	Verified bool
	Signer   string
	Diff     *mapnode.DiffResult
}

type CosignVerifyConfig struct {
	CertRef       string
	CertChain     string
	RekorURL      string
	OIDCIssuer    string
	RootCerts     *cryptox509.CertPool
	AllowInsecure bool
}

type BlobSignatureVerifier struct {
	annotations      map[string]string
	resourceRef      string
	annotationConfig AnnotationConfig
	identityList     []verificationIdentity
	CosignVerifyConfig
}

type BlobManifestFetcher struct {
	AnnotationConfig       AnnotationConfig
	resourceRefString      string
	ignoreFields           []string
	maxResourceManifestNum int
}

type ImageSignatureVerifier struct {
	resBundleRef         string
	onMemoryCacheEnabled bool
	annotationConfig     AnnotationConfig
	identityList         []verificationIdentity
	CosignVerifyConfig
}

func NewSignatureVerifier(objYAMLBytes []byte, sigRef string, pubkeyPath *string, signers []string, cosignVerifyConfig CosignVerifyConfig, annotationConfig AnnotationConfig) SignatureVerifier {
	panic("stub")
}

func (v *ImageSignatureVerifier) Verify() (bool, string, *int64, error) {
	panic("stub")
}

func (v *BlobSignatureVerifier) Verify() (bool, string, *int64, error) {
	panic("stub")
}

func NewManifestFetcher(resBundleRef, resourceRef string, annotationConfig AnnotationConfig, ignoreFields []string, maxResourceManifestNum int, allowInsecure bool) ManifestFetcher {
	panic("stub")
}

func (f *ImageManifestFetcher) Fetch(objYAMLBytes []byte) ([][]byte, string, error) {
	panic("stub")
}

func (f *ImageManifestFetcher) FetchAll() ([][]byte, error) {
	panic("stub")
}

func (f *BlobManifestFetcher) Fetch(objYAMLBytes []byte) ([][]byte, string, error) {
	panic("stub")
}

func (r *VerifyResult) String() string {
	panic("stub")
}

func GetConfigMapFromK8sObjectRef(objRef string) (interface{}, error) {
	panic("stub")
}

func VerifyManifest(objManifest []byte, vo *VerifyManifestOption) (*VerifyResult, error) {
	panic("stub")
}

const defaultDryRunNamespace = "default"

const imageCanonicalizePatternRegistry = "docker.io/"

const imageCanonicalizePatternTag = ":latest"

type VerifyResourceResult struct {
	Verified        bool
	InScope         bool
	Signer          string
	SignedTime      *time.Time
	SigRef          string
	Diff            *mapnode.DiffResult
	ContainerImages []kubeutil.ImageObject
	Provenances     []*Provenance
}

func (r *VerifyResourceResult) String() string {
	panic("stub")
}

func VerifyResource(obj interface{}, vo *VerifyResourceOption) (*VerifyResourceResult, error) {
	panic("stub")
}

type Embedme interface{}
