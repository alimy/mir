// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"strings"

	"github.com/alimy/mir/mirc/v5/templates"
	"github.com/spf13/cobra"
)

func init() {
	stylesCmd := &cobra.Command{
		Use:   "styles",
		Short: "list all suported style that engine temeplate code generate",
		Long:  "list all suported style that engine temeplate code generate",
		Run:   stylesRun,
	}
	register(stylesCmd)
}

// stylesRun run styles command
func stylesRun(_cmd *cobra.Command, _args []string) {
	fmt.Println(strings.Join(templates.Styles(), " "))
}
