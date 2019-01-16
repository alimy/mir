#!/usr/bin/env bash

# Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
# Use of this source code is governed by Apache License 2.0 that
# can be found in the LICENSE file.

set -o errexit
set -o nounset
set -o pipefail

function testModule() {
    for module_dir in $@; do
        cd ${module_dir}
        GO111MODULE=on go test -race ./...
        cd -
    done
}

testModule $@