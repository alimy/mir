package conf

import (
	"github.com/alimy/mir/v5/service"
	"github.com/fatih/color"
)

func setupMir() {
	service.SetOutput(color.Output)
}
