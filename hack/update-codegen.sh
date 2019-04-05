#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

source $(dirname $0)/../vendor/github.com/knative/test-infra/scripts/library.sh
CODEGEN_PKG=${CODEGEN_PKG:-$(cd ${REPO_ROOT_DIR}; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../../../k8s.io/code-generator)}
${CODEGEN_PKG}/generate-groups.sh "deepcopy,client,informer,lister" \
  github.com/nachocano/bitbucket-source/pkg/client github.com/nachocano/bitbucket-source/pkg/apis \
  "sources:v1alpha1" \
  --go-header-file ${REPO_ROOT_DIR}/hack/boilerplate/boilerplate.go.txt

${REPO_ROOT_DIR}/hack/update-deps.sh
