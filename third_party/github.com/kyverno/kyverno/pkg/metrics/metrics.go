package metrics

import (
	"context"
	"net/http"

	"github.com/kyverno/kyverno/pkg/config"

	kconfig "github.com/kyverno/kyverno/pkg/config"

	kyvernov1 "github.com/kyverno/kyverno/api/kyverno/v1"

	engineapi "github.com/kyverno/kyverno/pkg/engine/api"
)

type Recorder interface {
	Record(ClientQueryOperation)
	RecordWithContext(context.Context, ClientQueryOperation)
}

type clientQueryRecorder struct {
	manager MetricsConfigManager
	ns      string
	kind    string
	client  ClientType
}

func NamespacedClientQueryRecorder(m MetricsConfigManager, ns, kind string, client ClientType) Recorder {
	panic("stub")
}

func ClusteredClientQueryRecorder(m MetricsConfigManager, kind string, client ClientType) Recorder {
	panic("stub")
}

func (r *clientQueryRecorder) Record(clientQueryOperation ClientQueryOperation) {
	panic("stub")
}

func (r *clientQueryRecorder) RecordWithContext(ctx context.Context, clientQueryOperation ClientQueryOperation) {
	panic("stub")
}

const Enforce = "enforce"

const Audit = "audit"

const Cluster = "cluster"

const Namespaced = "namespaced"

const BackgroundTrue = "true"

const BackgroundFalse = "false"

const Validate = "validate"

const Mutate = "mutate"

const Generate = "generate"

const ImageVerify = "imageVerify"

const EmptyRuleType = "-"

const Pass = "pass"

const Fail = "fail"

const Warn = "warn"

const Error = "error"

const Skip = "skip"

const AdmissionRequest = "admission_request"

const BackgroundScan = "background_scan"

const ResourceCreated = "create"

const ResourceUpdated = "update"

const ResourceDeleted = "delete"

const ResourceConnected = "connect"

const ClientCreate = "create"

const ClientGet = "get"

const ClientList = "list"

const ClientUpdate = "update"

const ClientUpdateStatus = "update_status"

const ClientDelete = "delete"

const ClientDeleteCollection = "delete_collection"

const ClientWatch = "watch"

const ClientPatch = "patch"

const DynamicClient = "dynamic"

const KubeClient = "kubeclient"

const KyvernoClient = "kyverno"

const MetadataClient = "metadata"

const ApiServerClient = "apiserver"

const PolicyReportClient = "policyreport"

type ClientType string

type PolicyType string

type RuleExecutionCause string

type RuleResult string

type RuleType string

type PolicyBackgroundMode string

type ResourceRequestOperation string

type PolicyValidationMode string

type ClientQueryOperation string

func NewFakeMetricsConfig() *MetricsConfig {
	panic("stub")
}

func InitMetrics(ctx context.Context, disableMetricsExport bool, otel string, metricsAddr string, otelCollector string, metricsConfiguration config.MetricsConfiguration, transportCreds string, kubeClient interface{}, logger interface{}) (MetricsConfigManager, *http.ServeMux, interface{}, error) {
	panic("stub")
}

const MeterName = "kyverno"

type MetricsConfig struct {
	policyChangesMetric interface{}
	clientQueriesMetric interface{}
	config              kconfig.MetricsConfiguration
	Log                 interface{}
}

type MetricsConfigManager interface {
	Config() kconfig.MetricsConfiguration
	RecordPolicyChanges(ctx context.Context, policyValidationMode PolicyValidationMode, policyType PolicyType, policyBackgroundMode PolicyBackgroundMode, policyNamespace string, policyName string, policyChangeType string)
	RecordClientQueries(ctx context.Context, clientQueryOperation ClientQueryOperation, clientType ClientType, resourceKind string, resourceNamespace string)
}

func (m *MetricsConfig) Config() kconfig.MetricsConfiguration {
	panic("stub")
}

func ShutDownController(ctx context.Context, pusher interface{}) {
	panic("stub")
}

func NewOTLPGRPCConfig(ctx context.Context, endpoint string, certs string, kubeClient interface{}, log interface{}) (interface{}, error) {
	panic("stub")
}

func NewPrometheusConfig(ctx context.Context, log interface{}) (interface{}, *http.ServeMux, error) {
	panic("stub")
}

func (m *MetricsConfig) RecordPolicyChanges(ctx context.Context, policyValidationMode PolicyValidationMode, policyType PolicyType, policyBackgroundMode PolicyBackgroundMode, policyNamespace string, policyName string, policyChangeType string) {
	panic("stub")
}

func (m *MetricsConfig) RecordClientQueries(ctx context.Context, clientQueryOperation ClientQueryOperation, clientType ClientType, resourceKind string, resourceNamespace string) {
	panic("stub")
}

func ParsePolicyValidationMode(validationFailureAction kyvernov1.ValidationFailureAction) (PolicyValidationMode, error) {
	panic("stub")
}

func ParsePolicyBackgroundMode(policy kyvernov1.PolicyInterface) PolicyBackgroundMode {
	panic("stub")
}

func ParseRuleType(rule kyvernov1.Rule) RuleType {
	panic("stub")
}

func ParseResourceRequestOperation(requestOperationStr string) (ResourceRequestOperation, error) {
	return "", nil

}

func ParseRuleTypeFromEngineRuleResponse(rule engineapi.RuleResponse) RuleType {
	return ""

}

func GetPolicyInfos(policy kyvernov1.PolicyInterface) (string, string, PolicyType, PolicyBackgroundMode, PolicyValidationMode, error) {
	return "", "", "", "", "", nil

}

type Embedme interface{}
