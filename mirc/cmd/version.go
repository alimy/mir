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
