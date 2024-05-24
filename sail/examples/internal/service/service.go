// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package service

import (
	"net/http"

	sail "github.com/alimy/mir/sail/v4/service"
	"github.com/gin-gonic/gin"
)

func NewRuntime() sail.Runtime {
	webEngine, botEngine := newEngine(), newEngine()
	p := sail.NewHttpServerPool[*gin.Engine]()

	webServerAddr, botServerAddr := ":8080", ":8081"
	hsWeb := sail.NewBaseHttpService(p, webEngine, &http.Server{
		Addr: webServerAddr,
	})
	hsBot := sail.NewBaseHttpService(p, botEngine, &http.Server{
		Addr: botServerAddr,
	})

	webSrv := newWebService(hsWeb, webServerAddr)
	botSrv := newBotService(hsBot, botServerAddr)

	// init service
	sail.MustInitService(webSrv, botSrv)
	return p
}

func newEngine() *gin.Engine {
	e := gin.New()
	e.HandleMethodNotAllowed = true
	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	// 默认404
	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "Not Found",
		})
	})

	// 默认405
	e.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code": 405,
			"msg":  "Method Not Allowed",
		})
	})

	return e
}
