#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

dep ensure

rm -rf $(find vendor/ -name 'OWNERS')
rm -rf $(find vendor/ -name '*_test.go')
