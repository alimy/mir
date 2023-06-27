// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

//go:build generate
// +build generate

package main

import (
	"log"

	. "github.com/alimy/mir/v4/core"
	. "github.com/alimy/mir/v4/engine"

	_ "github.com/alimy/mir-example/v4/mirc/routes"
	_ "github.com/alimy/mir-example/v4/mirc/routes/v1"
	_ "github.com/alimy/mir-example/v4/mirc/routes/v2"
	_ "github.com/alimy/mir-example/v4/mirc/routes/v3"
)

//go:generate go run $GOFILE
func main() {
	log.Println("generate code start")
	opts := Options{
		RunMode(InSerialMode),
		GeneratorName(GeneratorGin),
		WatchCtxDone(true),
		SinkPath("auto"),
	}
	if err := Generate(opts); err != nil {
		log.Fatal(err)
	}
	log.Println("generate code finish")
}
