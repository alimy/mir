#!/usr/bin/env bash

# Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
# Use of this source code is governed by Apache License 2.0 that
# can be found in the LICENSE file.

set -o errexit
set -o nounset
set -o pipefail

root_dir=$(pwd)
echo "mode: count" > coverage.out
for module_dir in $@; do
    cd ${module_dir}
    GO111MODULE=on go test -v -race -covermode=atomic -coverprofile=profile.out ./... > tmp.out
	cat tmp.out
	if grep -q "^--- FAIL" tmp.out; then
		rm tmp.out
		exit 1
	elif grep -q "build failed" tmp.out; then
		rm tmp.out
		exit
	fi
	if [ -f profile.out ]; then
		cat profile.out | grep -v "mode:" >> ${root_dir}/coverage.out
		rm tmp.out
		rm profile.out
	fi
	cd -
done