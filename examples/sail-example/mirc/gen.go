// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

//go:build generate
// +build generate

package main

import (
	"log"

	. "github.com/alimy/mir/v5/core"
	. "github.com/alimy/mir/v5/engine"

	_ "github.com/alimy/mir/sail-example/v5/mirc/bot/v1"
	_ "github.com/alimy/mir/sail-example/v5/mirc/web"
	_ "github.com/alimy/mir/sail-example/v5/mirc/web/v1"
	_ "github.com/alimy/mir/sail-example/v5/mirc/web/v2"
	_ "github.com/alimy/mir/sail-example/v5/mirc/web/v3"
)

//go:generate go run $GOFILE
func main() {
	log.Println("generate code start")
	if err := Generate(
		UseGin(),
		Schema("web", "bot"),
		SinkPath("../auto"),
		UseRequestContext(),
		WatchCtxDone(true),
		RunMode(InSerialMode),
	); err != nil {
		log.Fatal(err)
	}
	log.Println("generate code finish")
}
