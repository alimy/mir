// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/alimy/mir/mirc/v4/internal/templates"
	"github.com/spf13/cobra"
)

var (
	dstPath    string
	pkgName    string
	style      []string
	mirPkgName string
	mirVersion string
	engVersion string
)

func init() {
	newCmd := &cobra.Command{
		Use:   "new",
		Short: "create template project",
		Long:  "create template project",
		Run:   newRun,
	}

	// parse flags for agentCmd
	newCmd.Flags().StringVarP(&dstPath, "dst", "d", ".", "genereted destination target directory")
	newCmd.Flags().StringVarP(&pkgName, "pkg", "p", "github.com/alimy/mir-example", "project's package name")
	newCmd.Flags().StringSliceVarP(&style, "style", "s", []string{"gin"}, "generated engine style eg: gin,chi,mux,hertz,echo,iris,fiber,macaron,httprouter")
	newCmd.Flags().StringVar(&mirPkgName, "mir", "", "mir replace package name or place")
	newCmd.Flags().StringVar(&mirVersion, "mir-version", "", "set mir version")
	newCmd.Flags().StringVar(&engVersion, "engine-version", "", "set engine version")

	// register agentCmd as sub-command
	register(newCmd)
}

// newRun run new command
func newRun(_cmd *cobra.Command, _args []string) {
	path, err := filepath.EvalSymlinks(dstPath)
	if err != nil {
		if os.IsNotExist(err) {
			if !filepath.IsAbs(dstPath) {
				cwd, err := os.Getwd()
				if err != nil {
					log.Fatal(err)
				}
				path = filepath.Join(cwd, dstPath)
			} else {
				path = dstPath
			}
		} else {
			log.Fatal(err)
		}
	}
	ctx := &templates.TmplCtx{
		PkgName:    pkgName,
		MirPkgName: mirPkgName,
		MirVersion: mirVersion,
		EngVersion: engVersion,
	}
	if err = genProject(ctx, path, style); err != nil {
		log.Fatal(err)
	}
}

func genProject(ctx *templates.TmplCtx, dstPath string, style []string) error {
	var (
		err               error
		filePath, dirPath string
		file              *os.File
	)

	t, err := templates.NewTemplate(style)
	if err != nil {
		return fmt.Errorf("not exist style for %s: %w", strings.Join(style, ":"), err)
	}

	for _, tmpl := range t.Templates() {
		filePath = filepath.Join(dstPath, tmpl.Name())
		dirPath = filepath.Dir(filePath)
		if err = os.MkdirAll(dirPath, 0755); err != nil {
			break
		}
		file, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			break
		}
		if err = tmpl.Execute(file, ctx); err != nil {
			break
		}
		file.Close()
	}
	return err
}
