// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"net/http"

	"github.com/gin-gonic/gin"

	api "github.com/alimy/mir/v3/examples/mirc/api/v1"
)

var _ api.Site = EmptySiteV1{}

// EmptySiteV1 implement api.Site interface
type EmptySiteV1 struct{}

func (EmptySiteV1) Chain() gin.HandlersChain {
	return nil
}

func (EmptySiteV1) Index(c *gin.Context) {
	c.String(http.StatusOK, "get index data (v1)")
}

func (EmptySiteV1) Articles(c *gin.Context) {
	c.String(http.StatusOK, "get articles data (v1)")
}
