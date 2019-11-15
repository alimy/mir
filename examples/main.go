// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/alimy/mir/v2/examples/mirc/gen/api"
	"github.com/alimy/mir/v2/examples/servants"
	"github.com/gin-gonic/gin"

	apiV1 "github.com/alimy/mir/v2/examples/mirc/gen/api/v1"
	apiV2 "github.com/alimy/mir/v2/examples/mirc/gen/api/v2"
)

func main() {
	e := gin.New()
	// register servant to engine
	api.RegisterSiteServant(e, servants.EmptySiteWithNoGroup{})
	apiV1.RegisterSiteServant(e, servants.EmptySiteV1{})
	apiV2.RegisterSiteServant(e, servants.EmptySiteV2{})
	// start servant service
	if err := e.Run(); err != nil {
		log.Fatal(err)
	}
}
