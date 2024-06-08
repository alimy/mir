// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

//go:build generate
// +build generate

package main

import (
	"log"

	. "github.com/alimy/mir/v4/core"
	. "github.com/alimy/mir/v4/engine"
	"github.com/gin-gonic/gin"

	_ "github.com/alimy/mir/sail-example/v4/mirc/bot/v1"
	_ "github.com/alimy/mir/sail-example/v4/mirc/web"
	_ "github.com/alimy/mir/sail-example/v4/mirc/web/v1"
	_ "github.com/alimy/mir/sail-example/v4/mirc/web/v2"
	_ "github.com/alimy/mir/sail-example/v4/mirc/web/v3"
)

//go:generate go run $GOFILE
func main() {
	log.Println("generate code start")
	opts := Options{
		UseGin(),
		UseRequestContext(),
		SinkPath("../auto"),
		WatchCtxDone(true),
		RunMode(InSerialMode),
		AssertType[*gin.Context](),
	}
	if err := Generate(opts); err != nil {
		log.Fatal(err)
	}
	log.Println("generate code finish")
}
