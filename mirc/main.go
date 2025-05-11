// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package main

import (
	"github.com/alimy/mir/mirc/v4/cmd"
)

func main() {
	// setup root cli command of application
	cmd.Setup(
		"mirc",             // command name
		"mir help toolkit", // command short describe
		"mir help tookit",  // command long describe
	)

	// execute start application
	cmd.Execute()
}
