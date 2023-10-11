package cli

import (
	"time"

	"github.com/sigstore/k8s-manifest-sigstore/pkg/k8smanifest"

	kubeutil "github.com/sigstore/k8s-manifest-sigstore/pkg/util/kubeutil"

	_ "embed"
)

func NewCmdApplyAfterVerify() interface{} {
	panic("stub")
}

const defaultChunkSize = 500

type KubectlOptions struct {
	ConfigFlags          interface{}
	PrintFlags           interface{}
	GetOptions           interface{}
	ApplyOptions         interface{}
	fieldManagerForApply string
}

func (o *KubectlOptions) SetKubeConfig(fpath, namespace string) {
	panic("stub")
}

func (o *KubectlOptions) InitGet(cmd interface{}) error {
	panic("stub")
}

func (o *KubectlOptions) Get(args []string, overrideNamespace string) ([]interface{}, error) {
	panic("stub")
}

func (o *KubectlOptions) InitApply(cmd interface{}, filename string) error {
	panic("stub")
}

func (o *KubectlOptions) Apply(filename string) error {
	panic("stub")
}

func NewCmdManifestBuild() interface{} {
	panic("stub")
}

const logLevelEnvKey = "K8S_MANIFEST_SIGSTORE_LOG_LEVEL"

var KOptions KubectlOptions

const filenameIfInputIsDir = "manifest.yaml"

func NewCmdSign() interface{} {
	panic("stub")
}

func NewCmdVerify() interface{} {
	panic("stub")
}

const resultAPIVersion = "cosign.sigstore.dev/v1alpha1"

const resultKind = "VerifyResourceResult"

const configTypeFile = "file"

const configTypeConstraint = "constraint"

const configTypeConfigMap = "configmap"

const defaultConfigKindForConfigMap = "ConfigMap"

const defaultConfigKindForConstraint = "ManifestIntegrityConstraint"

const defaultConfigFieldPathConstraint = "spec.parameters"

const defaultConfigFieldPathConfigMap = "data.\"config.yaml\""

const defaultMatchFieldPathConstraint = "spec.match"

const defaultInScopeObjectParameterFieldPathConstraint = "spec.parameters.inScopeObjects"

const defaultManifetBundleNamespace = "manifest-bundles"

type resourceResult struct {
	Object interface{}
	Result *k8smanifest.VerifyResourceResult
	Error  error
}

type provenanceSummary struct {
	Total     int
	Artifacts []string
}

type manifestResult struct {
	Name       string
	Signer     string
	SignedTime *time.Time
}

type provenanceResult struct {
	Summary provenanceSummary
	Items   []k8smanifest.Provenance
}

type summary struct {
	Total   int
	Valid   int
	Invalid int
}

type VerifyResourceResult struct {
	Embedme
	Summary    summary
	Manifests  []manifestResult
	Resources  []resourceResult
	Provenance *provenanceResult
}

type imageToBeUsed struct {
	kubeutil.ImageObject
	imageType k8smanifest.ArtifactType
}

func NewCmdVerifyResource() interface{} {
	panic("stub")
}

func (r resourceResult) MarshalJSON() ([]byte, error) {
	panic("stub")
}

func (r resourceResult) MarshalYAML() ([]byte, error) {
	panic("stub")
}

func NewVerifyResourceResult(results []resourceResult, provenanceEnabled bool) VerifyResourceResult {
	panic("stub")
}

func NewCmdVersion() interface{} {
	panic("stub")
}

type Embedme interface{}
