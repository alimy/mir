package main

import (
	"log"

	"github.com/alimy/mir/v2/core"
	"github.com/alimy/mir/v2/engine"
	"github.com/alimy/mir/v2/examples/mirc/routers"
	"github.com/gin-gonic/gin"
)

//go:generate go run main.go
func main() {
	log.Println("generate code start")
	entries := mirEntries()
	_ = engine.Generate(entries, &core.GenOpts{Name: core.GeneratorGin, OutPath: "./gen"})
	log.Println("generate code finish")
}

func mirEntries() []interface{} {
	return []interface{}{
		&routers.Site{Chain: gin.HandlersChain{gin.Logger()}},
	}
}
