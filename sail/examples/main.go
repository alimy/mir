// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/alimy/mir/sail/mir-example/v4/service"
	sail "github.com/alimy/mir/sail/v4/service"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/sourcegraph/conc"
)

func main() {
	webEngine, botEngine := service.NewEngine(), service.NewEngine()
	p := sail.NewHttpServerPool[*gin.Engine]()

	webServerAddr, botServerAddr := ":8080", ":8081"
	hsWeb := sail.NewBaseHttpService(p, webEngine, &http.Server{
		Addr: webServerAddr,
	})
	hsBot := sail.NewBaseHttpService(p, botEngine, &http.Server{
		Addr: botServerAddr,
	})

	webSrv := service.NewWebService(hsWeb, webServerAddr)
	botSrv := service.NewBotService(hsBot, botServerAddr)

	// init service
	sail.MustInitService(webSrv, botSrv)

	// start services
	wg := conc.NewWaitGroup()
	fmt.Fprintf(color.Output, "\nstarting run service...\n\n")
	p.Start(wg)

	// graceful stop services
	wg.Go(func() {
		quit := make(chan os.Signal, 1)
		// kill (no param) default send syscall.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		fmt.Fprintf(color.Output, "\nshutting down server...\n\n")
		p.Stop()
	})
	wg.Wait()
}
