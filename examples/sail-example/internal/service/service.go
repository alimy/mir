// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package service

import (
	"net/http"

	"github.com/alimy/mir/sail-example/v5/internal/conf"
	sail "github.com/alimy/mir/v5/service"
	"github.com/alimy/tryst/cfg"
	"github.com/gin-gonic/gin"
)

func NewRuntime() sail.Runtime {
	var ss []sail.Service
	p := sail.NewHttpServerPool[*gin.Engine]()

	// initial Web service
	cfg.Be("Web", func() {
		addr := conf.WebServerSetting.MyAddr()
		hs := sail.NewBaseHttpService(p, newEngine(), &http.Server{
			Addr:         addr,
			ReadTimeout:  conf.WebServerSetting.MyReadTimeout(),
			WriteTimeout: conf.WebServerSetting.MyWriteTimeout(),
		})
		ss = append(ss, newWebService(hs, addr))
	})

	// initial Bot service
	cfg.Be("Bot", func() {
		addr := conf.BotServerSetting.MyAddr()
		hs := sail.NewBaseHttpService(p, newEngine(), &http.Server{
			Addr:         addr,
			ReadTimeout:  conf.BotServerSetting.MyReadTimeout(),
			WriteTimeout: conf.BotServerSetting.MyWriteTimeout(),
		})
		ss = append(ss, newBotService(hs, addr))
	})

	// initial Docs service
	cfg.Be("Docs", func() {
		addr := conf.DocsServerSetting.MyAddr()
		hs := sail.NewBaseHttpService(p, newEngine(), &http.Server{
			Addr:         addr,
			ReadTimeout:  conf.DocsServerSetting.MyReadTimeout(),
			WriteTimeout: conf.DocsServerSetting.MyWriteTimeout(),
		})
		ss = append(ss, newDocsService(hs, addr))
	})

	// init service
	sail.MustInitService(ss...)
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
