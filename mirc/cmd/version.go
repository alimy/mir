// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/alimy/mir/mirc/v4/version"
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
	if version.BuildTime == "" || version.GitHash == "" {
		fmt.Printf("v%s\n", version.AppVer)
	} else {
		fmt.Printf("v%s\nBuildTime: %s\nGitHash: %s\n",
			version.AppVer, version.BuildTime, version.GitHash)
	}
}
