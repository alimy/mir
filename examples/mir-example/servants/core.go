// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"net/http"

	"github.com/alimy/mir/v5"
	"github.com/gin-gonic/gin"
)

type baseSrv struct{}

func (baseSrv) Bind(c *gin.Context, obj any) (err error) {
	if xerr := c.ShouldBind(obj); xerr != nil {
		err = mir.NewError(http.StatusBadRequest, xerr)
	}
	return
}

func (baseSrv) BindJson(c *gin.Context, obj any) (err error) {
	if xerr := c.BindJSON(obj); xerr != nil {
		err = mir.NewError(http.StatusBadRequest, xerr)
	}
	return
}

func (baseSrv) BindYaml(c *gin.Context, obj any) (err error) {
	if xerr := c.BindYAML(obj); xerr != nil {
		err = mir.NewError(http.StatusBadRequest, xerr)
	}
	return
}

func (baseSrv) Render(c *gin.Context, data any, err error) {
	if err == nil {
		c.JSON(http.StatusOK, data)
	} else if code, ok := mir.HttpStatusCode(err); ok {
		c.JSON(code, err.Error())
	} else {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}

func (baseSrv) RenderJson(c *gin.Context, data any, err error) {
	if err == nil {
		c.JSON(http.StatusOK, data)
	} else if code, ok := mir.HttpStatusCode(err); ok {
		c.JSON(code, err.Error())
	} else {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}

func (baseSrv) RenderJsonp(c *gin.Context, data any, err error) {
	if err == nil {
		c.JSONP(http.StatusOK, data)
	} else if code, ok := mir.HttpStatusCode(err); ok {
		c.JSONP(code, err.Error())
	} else {
		c.JSONP(http.StatusInternalServerError, err.Error())
	}
}

func (baseSrv) RenderYaml(c *gin.Context, data any, err error) {
	if err == nil {
		c.YAML(http.StatusOK, data)
	} else if code, ok := mir.HttpStatusCode(err); ok {
		c.YAML(code, err.Error())
	} else {
		c.YAML(http.StatusInternalServerError, err.Error())
	}
}

func (baseSrv) RenderXML(c *gin.Context, data any, err error) {
	if err == nil {
		c.XML(http.StatusOK, data)
	} else if code, ok := mir.HttpStatusCode(err); ok {
		c.XML(code, err.Error())
	} else {
		c.XML(http.StatusInternalServerError, err.Error())
	}
}
