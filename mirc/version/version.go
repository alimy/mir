// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package version

import "github.com/coreos/go-semver/semver"

var (
	// GitHash Value will be set during build
	GitHash = "Not provided"
	// BuildTime Value will be set during build
	BuildTime = "Not provided"
)

// AppVer version of Mirc
var AppVer = semver.Version{
	Major:      2,
	Minor:      1,
	Patch:      0,
}
