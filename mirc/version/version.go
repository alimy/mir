// Copyright 2025 Michael Li <alimy@niubiu.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package version

import (
	"debug/buildinfo"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/Masterminds/semver/v3"
)

var (
	// GitHash Value will be set during build
	GitHash = ""

	// BuildTime Value will be set during build
	BuildTime = sync.OnceValue(getBuildTime)

	// AppVer version of Mirc
	AppVer = semver.MustParse("v5.2.0")
)

func ShowInfo() string {
	if GitHash == "" {
		return fmt.Sprintf("mirc v%s build(%s)\n", AppVer, BuildTime())
	} else {
		return fmt.Sprintf("mirc v%s build(%s %s)\n", AppVer, GitHash, BuildTime())
	}
}

func getBuildTime() (buildTime string) {
	buildTime = time.Now().Local().Format(time.DateTime)
	exe, err := os.Executable()
	if err != nil {
		return
	}
	info, err := buildinfo.ReadFile(exe)
	if err != nil {
		return
	}
	for _, s := range info.Settings {
		if s.Key == "vcs.time" && s.Value != "" {
			if t, err := time.Parse(time.RFC3339, s.Value); err == nil {
				buildTime = t.Local().Format(time.DateTime)
			}
			break
		}
	}
	return
}
