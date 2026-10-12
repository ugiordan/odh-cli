# NetworkPolicy guidance for OpenShift AI (ODH)

This document describes the recommended NetworkPolicy posture for clusters running OpenShift AI components managed by the odh-cli tooling.

Why NetworkPolicy?

- NetworkPolicy provides pod-level network isolation which complements TLS and authentication checks already present in odh-cli.
- During upgrades and migrations the CLI may remove/add components: having explicit network policies reduces blast radius and helps ensure services are only reachable by intended callers.

Recommendations

1. Default-deny for sensitive namespaces

Create a "default-deny" NetworkPolicy in namespaces hosting OpenShift AI components (for example: `opendatahub`, `odh`, or any project where workbenches/run-time pods live). Example:

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: default-deny-ingress
spec:
  podSelector: {}
  policyTypes:
  - Ingress
  ingress: []
```

This denies all ingress to pods in the namespace by default; explicit allow policies should then be added for traffic that needs to flow (operator, control-plane, ingress controllers, etc.).

2. Fine-grained allow policies

For services that require external access (ingress routes, APIs), create narrow NetworkPolicies that allow traffic only from the required source pods, namespaces, or CIDR ranges.

3. Avoid overly permissive selectors

Watch out for policies like `podSelector: {}` combined with wide `from` blocks which effectively allow broad access. Prefer explicit selectors and namespaceSelectors where appropriate.

Next steps for odh-cli

- Implement lint checks that detect the absence of default-deny policies in relevant namespaces and warn/remediate during pre-upgrade checks.
- Optionally provide a migration action that can create baseline NetworkPolicies (with dry-run and confirmation prompts) to help clusters reach the expected security posture.
