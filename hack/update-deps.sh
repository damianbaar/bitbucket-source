#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

source $(dirname $0)/../vendor/github.com/knative/test-infra/scripts/library.sh
cd ${REPO_ROOT_DIR}
dep ensure

rm -rf $(find vendor/ -name 'BUILD')
rm -rf $(find vendor/ -name 'BUILD.bazel')
