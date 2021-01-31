// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"net/http"

	"github.com/gin-gonic/gin"

	api "github.com/alimy/mir/v2/examples/mirc/auto/api/v1"
)

type emptySiteV1 struct {
	api.UnimplementedSiteServant
}

func (*emptySiteV1) Index(c *gin.Context) {
	c.String(http.StatusOK, "get index data (v1)")
}

func (*emptySiteV1) Articles(c *gin.Context) {
	c.String(http.StatusOK, "get articles data (v1)")
}

func newSiteV1Srv() api.Site {
	return &emptySiteV1{}
}
