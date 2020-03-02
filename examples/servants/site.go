// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"net/http"

	"github.com/alimy/mir/v2/examples/mirc/gen/api"
	"github.com/gin-gonic/gin"
)

var _ api.Site = EmptySiteWithNoGroup{}

// EmptySiteWithNoGroup implement api.Site interface
type EmptySiteWithNoGroup struct{}

func (EmptySiteWithNoGroup) Chain() gin.HandlersChain {
	return gin.HandlersChain{gin.Logger()}
}

func (EmptySiteWithNoGroup) Index(c *gin.Context) {
	c.String(http.StatusOK, "get index data (v1)")
}

func (EmptySiteWithNoGroup) Articles(c *gin.Context) {
	c.String(http.StatusOK, "get articles data (v1)")
}
