// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/alimy/mir/mirc/v2/version"
	"github.com/spf13/cobra"
)

func init() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "show version information",
		Long:  "show version information",
		Run:   versionRun,
	}
	register(versionCmd)
}

func versionRun(_cmd *cobra.Command, _args []string) {
	fmt.Printf("v%s\nBuildTime:%s\nBuildGitSHA:%s\n",
		version.MircVer, version.BuildTime, version.GitHash)
}
