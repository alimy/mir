// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package base

import (
	"net/http"

	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
)

type BaseSrv struct{}

func (BaseSrv) Bind(c *gin.Context, obj any) (err mir.Error) {
	if xerr := c.ShouldBind(obj); xerr != nil {
		err = mir.NewError(http.StatusBadRequest, xerr)
	}
	return
}

func (BaseSrv) BindJson(c *gin.Context, obj any) (err mir.Error) {
	if xerr := c.BindJSON(obj); xerr != nil {
		err = mir.NewError(http.StatusBadRequest, xerr)
	}
	return
}

func (BaseSrv) BindYaml(c *gin.Context, obj any) (err mir.Error) {
	if xerr := c.BindYAML(obj); xerr != nil {
		err = mir.NewError(http.StatusBadRequest, xerr)
	}
	return
}

func (BaseSrv) Render(c *gin.Context, data any, err mir.Error) {
	if err == nil {
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(err.StatusCode(), err.Error())
	}
}

func (BaseSrv) RenderJson(c *gin.Context, data any, err mir.Error) {
	if err == nil {
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(err.StatusCode(), err.Error())
	}
}

func (BaseSrv) RenderJsonp(c *gin.Context, data any, err mir.Error) {
	if err == nil {
		c.JSONP(http.StatusOK, data)
	} else {
		c.JSONP(err.StatusCode(), err.Error())
	}
}

func (BaseSrv) RenderYaml(c *gin.Context, data any, err mir.Error) {
	if err == nil {
		c.YAML(http.StatusOK, data)
	} else {
		c.YAML(err.StatusCode(), err.Error())
	}
}

func (BaseSrv) RenderXML(c *gin.Context, data any, err mir.Error) {
	if err == nil {
		c.XML(http.StatusOK, data)
	} else {
		c.XML(err.StatusCode(), err.Error())
	}
}
