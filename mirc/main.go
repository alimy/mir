package main

import (
	"github.com/alimy/mir/mirc/v5/cmd"
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
