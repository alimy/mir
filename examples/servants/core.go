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

func (baseSrv) Bind(c *gin.Context, obj any) (err mir.Error) {
	if xerr := c.ShouldBind(obj); xerr != nil {
		err = mir.NewError(http.StatusBadRequest, xerr)
	}
	return
}

func (baseSrv) Render(c *gin.Context, data any, err mir.Error) {
	if err == nil {
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(err.StatusCode(), err.Error())
	}
}

func (baseSrv) BindByName(name string, c *gin.Context, obj any) (err mir.Error) {
	var xerr error
	switch name {
	case "yaml":
		xerr = c.BindYAML(obj)
	case "json":
		xerr = c.BindJSON(obj)
	default:
		xerr = c.ShouldBind(obj)
	}
	if xerr != nil {
		err = mir.NewError(http.StatusBadRequest, xerr)
	}
	return
}

func (baseSrv) RenderByName(name string, c *gin.Context, data any, err mir.Error) {
	switch name {
	case "jsonp":
		if err == nil {
			c.JSONP(http.StatusOK, data)
		} else {
			c.JSONP(err.StatusCode(), err.Error())
		}
	case "json":
		fallthrough
	default:
		if err == nil {
			c.JSON(http.StatusOK, data)
		} else {
			c.JSON(err.StatusCode(), err.Error())
		}
	}
}
