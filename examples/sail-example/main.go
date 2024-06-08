// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package main

import (
	"github.com/alimy/mir/sail-example/v4/cmd"
	_ "github.com/alimy/mir/sail-example/v4/cmd/migrate"
	_ "github.com/alimy/mir/sail-example/v4/cmd/serve"
)

func main() {
	cmd.Execute()
}
