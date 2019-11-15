package main

import (
	"github.com/alimy/mir/v2/core"
	"github.com/alimy/mir/v2/engine"
	"github.com/alimy/mir/v2/examples/mirc/routers"
	"github.com/gin-gonic/gin"
)

//go:generate go run main.go
func main() {
	entries := mirEntries()
	_ = engine.Generate(entries, &core.GenOpts{Name: core.GeneratorGin, OutPath: "./gen"})
}

func mirEntries() []interface{} {
	return []interface{}{
		&routers.Site{Chain: gin.HandlersChain{gin.Logger()}},
	}
}
