// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"fmt"
)

const (
	series = "v0.1-dev"
)

var (
	version   = "unknown"
	commitID  = "unknown"
	buildDate = "unknown"
	buildTags = "unknown"
)

type BuildInfo struct {
	Series    string `json:"series"`
	Version   string `json:"version"`
	Sum       string `json:"sum"`
	BuildDate string `json:"build_date"`
	BuildTags string `json:"build_tags"`
}

func VersionInfo() string {
	return fmt.Sprintf("mir-example %s (build:%s %s)", version, commitID, buildDate)
}

func ReadBuildInfo() *BuildInfo {
	return &BuildInfo{
		Series:    series,
		Version:   version,
		Sum:       commitID,
		BuildDate: buildDate,
		BuildTags: buildTags,
	}
}
