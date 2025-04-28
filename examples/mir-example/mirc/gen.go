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
)

//go:generate go run $GOFILE
func main() {
	log.Println("generate code start")
	if err := Generate(
		Schema("./routes"),
		UseGin(),
		UseRequestContext(),
		SinkPath("auto"),
		WatchCtxDone(true),
		RunMode(InSerialMode),
	); err != nil {
		log.Fatal(err)
	}
	log.Println("generate code finish")
}
