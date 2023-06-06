#!/usr/bin/env bats

@test "[flux policy] reject because fluxcd value is not allowed" {
  run kwctl run annotated-policy.wasm \
    -r test_data/requests/pod_creation_flux_cat.json \
    --settings-path test_data/settings/flux-policy.yaml

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request rejected
  [ "$status" -eq 0 ]
  [ $(expr "$output" : '.*allowed.*false') -ne 0 ]
  [ $(expr "$output" : ".*The only approved FluxCD annotations are.*") -ne 0 ]
}

@test "[flux policy] accept because fluxcd value is allowed" {
  run kwctl run annotated-policy.wasm \
    -r test_data/requests/pod_creation_flux_cow.json \
    --settings-path test_data/settings/flux-policy.yaml

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request accepted
  [ "$status" -eq 0 ]
  [ $(expr "$output" : '.*allowed.*true') -ne 0 ]
}
