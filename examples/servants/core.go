// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"net/http"

	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
)

type baseSrv struct{}

func (baseSrv) Bind(c *gin.Context, obj any) mir.Error {
	if err := c.ShouldBind(obj); err != nil {
		mir.NewError(http.StatusBadRequest, err)
	}
	return nil
}

func (baseSrv) Render(c *gin.Context, data any, err mir.Error) {
	if err == nil {
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(err.StatusCode(), err.Error())
	}
}
