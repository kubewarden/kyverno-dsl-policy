package policychanges

import (
	"context"

	kyvernov1 "github.com/kyverno/kyverno/api/kyverno/v1"
	"github.com/kyverno/kyverno/pkg/metrics"
)

func RegisterPolicy(ctx context.Context, m metrics.MetricsConfigManager, policy kyvernov1.PolicyInterface, policyChangeType PolicyChangeType) error {
	panic("stub")
}

const PolicyCreated = "created"

const PolicyUpdated = "updated"

const PolicyDeleted = "deleted"

type PolicyChangeType string

type Embedme interface{}
