// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"net/http"

	"github.com/alimy/mir/v2/examples/mirc/auto/api"
	"github.com/gin-gonic/gin"
)

type emptySiteWithNoGroup struct {
	api.UnimplementedSiteServant
}

func (emptySiteWithNoGroup) Chain() gin.HandlersChain {
	return gin.HandlersChain{gin.Logger()}
}

func (emptySiteWithNoGroup) Index(c *gin.Context) {
	c.String(http.StatusOK, "get index data")
}

func (emptySiteWithNoGroup) Articles(c *gin.Context) {
	c.String(http.StatusOK, "get articles data")
}

func newSiteSrv() api.Site {
	return &emptySiteWithNoGroup{}
}
