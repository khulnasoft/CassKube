#!/bin/bash

set -euf -o pipefail
set -x

# uses helm v3.1.2, kubectl v1.18.2, yq v3.2.1

bundle="$(mktemp)"

k8sVer="$(kubectl version --short | grep Server | egrep -o 'v[0-9].[0-9]+')"

echo '---' >> "$bundle"
cat operator/deploy/namespace.yaml | yq r - >> "$bundle"

echo '---' >> "$bundle"
helm template ./charts/casskube-chart/ -n casskube --validate=true | kubectl create --dry-run=client -o yaml -n casskube -f - >> "$bundle"

mv "$bundle" docs/user/casskube-manifests-$k8sVer.yaml
