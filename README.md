> **WARNING:** this is an experimental policy

This policy aims to provide a way to reuse Kyverno policies with Kubewarden.

## Policy goals

The policy aims to allow the creation of validating and mutating policies.
It does not intend to provide a way to use Kyverno's generating policies.

## Known limitations

Technical limitations caused by Go compiler not having a mature
[WASI](https://wasi.dev/) support:

* The policy requires Go 1.21 or later. Currently this is not yet published,
  hence a Go compiler built from the [`master`](https://github.com/golang/go)
  is required
* The size of the policy is huge: approximatively 90 Mb
* The 1st run of the policy is slow because code has to be optimized and compiled
  by Wasmtime for the local machine. Subsequent invocations are fast
* This policy requires Kubewarden to support the new `wasi` execution mode. This
  mode provides slower evaluation time compared to the traditional `wapc` one.
  Once [this](https://github.com/golang/go/issues/42372) Go issue is addressed, the
  policy will be rewritten to make use of the traditional Kubewarden policy
  interface.

Limitations due to the maturity of this code base:

* Policy settings are not supported
* Mutation support is not yet done
* Context aware information are not provided to the policy

## Usage

The policy provides the Kyverno evaluation engine, the actual policy has to
be provided via the policy settings.

This is done using the `policy` key, that has as value a string holding the actual
YAML definition of the Kyverno policy.

For example, to deploy [this](https://kyverno.io/policies/flux/verify-flux-sources/verify-flux-sources/) Kyverno policy,
the settings of the policy should look like that:

```yaml
policy: |
  apiVersion: kyverno.io/v1
  kind: ClusterPolicy
  metadata:
    name: allowed-annotations
    annotations:
      policies.kyverno.io/title: Allowed Annotations
      policies.kyverno.io/category: Other
      policies.kyverno.io/severity: medium
      kyverno.io/kyverno-version: 1.6.0
      policies.kyverno.io/minversion: 1.6.0
      kyverno.io/kubernetes-version: "1.23"
      policies.kyverno.io/subject: Pod, Annotation
      policies.kyverno.io/description: >-
        Rather than creating a deny list of annotations, it may be more useful
        to invert that list and create an allow list which then denies any others.
        This policy demonstrates how to allow two annotations with a specific key
        name of fluxcd.io/ while denying others that do not meet the pattern.      
  spec:
    validationFailureAction: audit
    background: true
    rules:
    - name: allowed-fluxcd-annotations
      match:
        any:
        - resources:
            kinds:
              - Pod
      validate:
        message: The only approved FluxCD annotations are `fluxcd.io/cow` and `fluxcd.io/dog`.
        deny:
          conditions:
            all:
            - key: "{{ request.object.metadata.annotations.keys(@)[?contains(@, 'fluxcd.io/')] }}"
              operator: AnyNotIn
              value:
              - fluxcd.io/cow
              - fluxcd.io/dog
```

> **Note:** this is just a copy & paste of the official Kyverno policy

This means the actual Kubewarden policy deployment would look like that:

```yaml
apiVersion: policies.kubewarden.io/v1
kind: ClusterAdmissionPolicy
metadata:
  name: allowed-flux-annotations
spec:
  policyServer: reserved-instance-for-tenant-a
  module: registry://ghcr.io/kubewarden/policies/kyverno-dsl:v0.1.0
  rules:
  - apiGroups: [""]
    apiVersions: ["v1"]
    resources: ["pods"]
    operations:
    - CREATE
    - UPDATE
  mutating: false
  settings:
    policy: |
      apiVersion: kyverno.io/v1
      kind: ClusterPolicy
      metadata:
        name: allowed-annotations
        annotations:
          policies.kyverno.io/title: Allowed Annotations
          policies.kyverno.io/category: Other
          policies.kyverno.io/severity: medium
          kyverno.io/kyverno-version: 1.6.0
          policies.kyverno.io/minversion: 1.6.0
          kyverno.io/kubernetes-version: "1.23"
          policies.kyverno.io/subject: Pod, Annotation
          policies.kyverno.io/description: >-
            Rather than creating a deny list of annotations, it may be more useful
            to invert that list and create an allow list which then denies any others.
            This policy demonstrates how to allow two annotations with a specific key
            name of fluxcd.io/ while denying others that do not meet the pattern.      
      spec:
        validationFailureAction: audit
        background: true
        rules:
        - name: allowed-fluxcd-annotations
          match:
            any:
            - resources:
                kinds:
                  - Pod
          validate:
            message: The only approved FluxCD annotations are `fluxcd.io/cow` and `fluxcd.io/dog`.
            deny:
              conditions:
                all:
                - key: "{{ request.object.metadata.annotations.keys(@)[?contains(@, 'fluxcd.io/')] }}"
                  operator: AnyNotIn
                  value:
                  - fluxcd.io/cow
                  - fluxcd.io/dog
```
