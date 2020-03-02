// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"net/http"

	"github.com/gin-gonic/gin"

	api "github.com/alimy/mir/v2/examples/mirc/gen/api/v2"
)

var _ api.Site = EmptySiteV2{}

// EmptySiteV2 implement api.Site interface
type EmptySiteV2 struct{}

func (EmptySiteV2) Index(c *gin.Context) {
	c.String(http.StatusOK, "get index data (v2)")
}

func (EmptySiteV2) Articles(c *gin.Context) {
	c.String(http.StatusOK, "get articles data (v2)")
}

func (EmptySiteV2) Category(c *gin.Context) {
	c.String(http.StatusOK, "get Category data (v2)")
}
