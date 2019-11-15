// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/alimy/mir/v2/core"
	"github.com/alimy/mir/v2/engine"
	"github.com/alimy/mir/v2/examples/mirc/routers/v1"
	"github.com/alimy/mir/v2/examples/mirc/routers/v2"
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
		&v1.Site{Chain: gin.HandlersChain{gin.Logger()}},
		&v2.Site{Chain: gin.HandlersChain{gin.Logger()}},
	}
}
