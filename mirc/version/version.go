// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package version

import (
	"github.com/Masterminds/semver/v3"
)

var (
	// GitHash Value will be set during build
	GitHash = ""

	// BuildTime Value will be set during build
	BuildTime = ""

	// AppVer version of Mirc
	AppVer = semver.MustParse("v2.6.3")
)
