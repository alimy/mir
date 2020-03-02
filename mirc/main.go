package main

import (
	"github.com/alimy/mir/mirc/v2/cmd"
)

func main() {
	// setup root cli command of application
	cmd.Setup(
		"mir",              // command name
		"mir help toolkit", // command short describe
		"mir help tookit",  // command long describe
	)

	// execute start application
	cmd.Execute()
}
