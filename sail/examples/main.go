// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/alimy/mir/sail/mir-example/v4/service"
	"github.com/fatih/color"
	"github.com/sourcegraph/conc"
)

func main() {
	runtime := service.NewRuntime()

	// start services
	wg := conc.NewWaitGroup()
	fmt.Fprintf(color.Output, "\nstarting run service...\n\n")
	runtime.Start(wg)

	// graceful stop services
	wg.Go(func() {
		quit := make(chan os.Signal, 1)
		// kill (no param) default send syscall.SIGTERM
		// kill -2 is syscall.SIGINT
		// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		fmt.Fprintf(color.Output, "\nshutting down server...\n\n")
		runtime.Stop()
	})
	wg.Wait()
}
