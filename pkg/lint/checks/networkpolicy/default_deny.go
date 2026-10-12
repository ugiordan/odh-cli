package networkpolicy

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/opendatahub-io/odh-cli/pkg/lint/check"
	"github.com/opendatahub-io/odh-cli/pkg/lint/result"
)

// DefaultDenyCheck verifies that namespaces hosting OpenShift AI components have
// a default-deny NetworkPolicy in place. This file contains a minimal, reviewable
// implementation to be extended in follow-up work.
type DefaultDenyCheck struct {
	check.BaseCheck
}

func NewDefaultDenyCheck() *DefaultDenyCheck {
	return &DefaultDenyCheck{
		BaseCheck: check.BaseCheck{
			CheckGroup:       check.GroupSecurity,
			Kind:             "NetworkPolicy",
			Type:             check.CheckTypeConfiguration,
			CheckID:          "security.networkpolicy.default-deny",
			CheckName:        "Security :: NetworkPolicy :: Default Deny",
			CheckDescription: "Detects whether namespaces hosting OpenShift AI components have a default-deny NetworkPolicy.",
			CheckRemediation: "Create a NetworkPolicy that denies all ingress by default and explicitly allows required traffic between components.",
		},
	}
}

func (c *DefaultDenyCheck) CanApply(ctx context.Context, target check.Target) (bool, error) {
	// Conservative default: this check can apply in any environment.
	return true, nil
}

func (c *DefaultDenyCheck) Validate(ctx context.Context, target check.Target) (*result.DiagnosticResult, error) {
	dr := result.NewDiagnosticResult(c)

	// TODO: implement actual discovery of NetworkPolicy resources and validation logic.
	// For this branch we add a clear placeholder condition so reviewers can see where
	// the check will hook into the linting framework. Follow-up patch will implement
	// proper discovery using the dynamic client and resources.ResourceType for
	// NetworkPolicy.
	dr.SetCondition(check.NewCondition(
		check.ConditionTypeConfigured,
		metav1.ConditionFalse,
		check.WithReason(check.ReasonConfigurationInvalid),
		check.WithMessage("NetworkPolicy validation not yet implemented in this branch"),
		check.WithImpact(result.ImpactAdvisory),
	))

	return dr, nil
}
