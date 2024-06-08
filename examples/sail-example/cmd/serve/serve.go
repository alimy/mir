// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package serve

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alimy/mir/sail-example/v4/cmd"
	"github.com/alimy/mir/sail-example/v4/internal/conf"
	"github.com/alimy/mir/sail-example/v4/internal/service"

	"github.com/fatih/color"
	"github.com/sourcegraph/conc"
	"github.com/spf13/cobra"
	"go.uber.org/automaxprocs/maxprocs"
)

var (
	noDefaultFeatures bool
	features          []string
)

func init() {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "start sail examples server",
		Long:  "start sail examples server",
		Run:   serveRun,
	}

	serveCmd.Flags().BoolVar(&noDefaultFeatures, "no-default-features", false, "whether not use default features")
	serveCmd.Flags().StringSliceVarP(&features, "features", "f", []string{}, "use special features")

	cmd.Register(serveCmd)
}

func serveRun(_cmd *cobra.Command, _args []string) {
	// set maxprocs automatic
	maxprocs.Set(maxprocs.Logger(log.Printf))

	// initial configure
	conf.Initial(features, noDefaultFeatures)

	// create service runtime
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
