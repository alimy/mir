// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/alimy/mir/v2/examples/servants"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.New()

	// register servants to engine
	servants.RegisterServants(e)

	// start servant service
	if err := e.Run(); err != nil {
		log.Fatal(err)
	}
}
