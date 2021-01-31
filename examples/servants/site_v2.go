// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"net/http"

	"github.com/gin-gonic/gin"

	api "github.com/alimy/mir/v2/examples/mirc/auto/api/v2"
)

type emptySiteV2 struct {
	api.UnimplementedSiteServant
}

func (*emptySiteV2) Chain() gin.HandlersChain {
	return gin.HandlersChain{gin.Logger()}
}

func (*emptySiteV2) Index(c *gin.Context) {
	c.String(http.StatusOK, "get index data (v2)")
}

func (*emptySiteV2) Articles(c *gin.Context) {
	c.String(http.StatusOK, "get articles data (v2)")
}

func (*emptySiteV2) Category(c *gin.Context) {
	c.String(http.StatusOK, "get Category data (v2)")
}

func newSiteV2Srv() api.Site {
	return &emptySiteV2{}
}
