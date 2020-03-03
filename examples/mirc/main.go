// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/alimy/mir/v2/core"
	"github.com/alimy/mir/v2/engine"

	routes "github.com/alimy/mir/v2/examples/mirc/routes"
	v1 "github.com/alimy/mir/v2/examples/mirc/routes/v1"
	v2 "github.com/alimy/mir/v2/examples/mirc/routes/v2"
)

//go:generate go run main.go

func main() {
	log.Println("generate code start")
	entries := mirEntries()
	opts := &core.Options{
		GeneratorName: core.GeneratorGin,
		OutPath:       "./gen"}
	if err := engine.Generate(entries, opts); err != nil {
		log.Fatal(err)
	}
	log.Println("generate code finish")
}

func mirEntries() []interface{} {
	return []interface{}{
		new(routes.Site),
		new(v1.Site),
		new(v2.Site),
	}
}
