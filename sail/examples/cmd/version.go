// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/alimy/mir/sail/examples/v4/internal/conf"
	"github.com/spf13/cobra"
)

func init() {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "show version information",
		Long:  "show version information",
		Run:   versionRun,
	}
	Register(versionCmd)
}

func versionRun(_cmd *cobra.Command, _args []string) {
	fmt.Println(conf.VersionInfo())
}
