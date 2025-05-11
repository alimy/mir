// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package migrate

import (
	"fmt"

	"github.com/alimy/mir/sail-example/v5/cmd"
	"github.com/spf13/cobra"
)

func init() {
	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "migrate database data",
		Long:  "miegrate database data when mir-examples upgrade",
		Run:   migrateRun,
	}
	cmd.Register(migrateCmd)
}

func migrateRun(_cmd *cobra.Command, _args []string) {
	// TODO: add some logic for migrate cmd feature
	fmt.Println("sorry, this feature is not implemented yet.")
}
