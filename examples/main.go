// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/alimy/mir/v3/examples/mirc/api"
	"github.com/alimy/mir/v3/examples/mirc/api/v1"
	"github.com/alimy/mir/v3/examples/mirc/api/v2"
	"github.com/alimy/mir/v3/examples/servants"
)

func main() {
	e := gin.New()

	// register servants to engine
	registerServants(e)

	// start servant service
	if err := e.Run(); err != nil {
		log.Fatal(err)
	}
}

func registerServants(e *gin.Engine) {
	// register default group routes
	api.RegisterSiteServant(e, servants.EmptySiteWithNoGroup{})

	// register routes for group v1
	v1.RegisterSiteServant(e, servants.EmptySiteV1{})

	// register routes for group v2
	v2.RegisterSiteServant(e, servants.EmptySiteV2{})
}
