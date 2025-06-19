// Copyright 2025 Michael Li <alimy@niubiu.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/alimy/mir/mirc/v5/version"
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
	if version.GitHash == "" {
		fmt.Printf("mirc v%s build(%s)\n", version.AppVer, version.BuildTime())
	} else {
		fmt.Printf("mirc v%s build(%s %s)\n", version.AppVer, version.GitHash, version.BuildTime())
	}
}
