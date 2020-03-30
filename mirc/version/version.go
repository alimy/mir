// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package version

import "github.com/coreos/go-semver/semver"

var (
	// GitHash Value will be set during build
	GitHash = ""
	// BuildTime Value will be set during build
	BuildTime = ""
)

// AppVer version of Mirc
var AppVer = semver.Version{
	Major: 2,
	Minor: 3,
	Patch: 2,
}
