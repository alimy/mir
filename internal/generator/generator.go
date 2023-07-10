// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package generator

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"sync"
	"text/template"

	"github.com/alimy/mir/v4/core"
	"github.com/alimy/mir/v4/internal/naming"
	"github.com/alimy/mir/v4/internal/utils"
)

func init() {
	core.RegisterGenerators(
		&mirGenerator{name: core.GeneratorGin},
		&mirGenerator{name: core.GeneratorChi},
		&mirGenerator{name: core.GeneratorMux},
		&mirGenerator{name: core.GeneratorHertz},
		&mirGenerator{name: core.GeneratorEcho},
		&mirGenerator{name: core.GeneratorIris},
		&mirGenerator{name: core.GeneratorFiber},
		&mirGenerator{name: core.GeneratorMacaron},
		&mirGenerator{name: core.GeneratorHttpRouter},
	)
}

type mirGenerator struct {
	sinkPath  string
	name      string
	isCleanup bool
}

type mirWriter struct {
	ns   naming.NamingStrategy
	tmpl *template.Template
}

func (w *mirWriter) Write(dirPath string, iface *core.IfaceDescriptor) error {
	fileName := w.ns.Naming(iface.TypeName) + ".go"
	filePath := filepath.Join(dirPath, fileName)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err == nil {
		defer func() {
			_ = file.Close()
		}()
		if err = w.tmpl.Execute(file, iface); err == nil {
			core.Logus("generated iface: %s.%s to file: %s", iface.PkgName, iface.TypeName, filePath)
		}
	}
	return err
}

// Name name of generator
func (g *mirGenerator) Name() string {
	return g.name
}

// Init init generator
func (g *mirGenerator) Init(opts *core.GeneratorOpts) (err error) {
	if opts == nil {
		return errors.New("init opts is nil")
	}
	g.isCleanup = opts.Cleanup
	g.sinkPath, err = evalSinkPath(opts.SinkPath)
	return
}

// Generate serial generate interface code
func (g *mirGenerator) Generate(ds core.Descriptors) error {
	// cleanup out first if need
	g.cleanup()
	return generate(g.name, g.sinkPath, ds)
}

// GenerateContext concurrent generate interface code
func (g *mirGenerator) GenerateContext(ctx core.MirCtx) {
	tmpl, err := templateFrom(g.name)
	if err != nil {
		ctx.Cancel(err)
		return
	}
	apiPath := filepath.Join(g.sinkPath, "api")
	ifaceSource, _ := ctx.Pipe()
	onceSet := utils.NewOnceSet(func(path string) error {
		return os.MkdirAll(path, 0755)
	})

	// cleanup out first if need
	g.cleanup()

	var t *template.Template
	wg := &sync.WaitGroup{}
	ns := naming.NewSnakeNamingStrategy()
	inOutsMap := make(map[string]utils.Set)
	for iface := range ifaceSource {
		dirPath := filepath.Join(apiPath, iface.Group)
		if err = onceSet.Add(dirPath); err != nil {
			goto FuckErr
		}
		if t, err = tmpl.Clone(); err != nil {
			goto FuckErr
		}

		// setup inOuts for IfaceDescriptor
		filter, exist := inOutsMap[iface.Group]
		if !exist {
			filter = utils.NewStrSet()
			inOutsMap[iface.Group] = filter
		}
		var inouts []reflect.Type
		for _, typ := range iface.AllInOuts() {
			if typ.PkgPath() == iface.PkgPath {
				if err := filter.Add(typ.Name()); err == nil {
					inouts = append(inouts, typ)
				}
			} else {
				inouts = append(inouts, typ)
			}
		}
		iface.SetInnerInOuts(inouts)

		writer := &mirWriter{tmpl: t, ns: ns}
		wg.Add(1)
		go func(ctx core.MirCtx, wg *sync.WaitGroup, writer *mirWriter, iface *core.IfaceDescriptor) {
			defer wg.Done()

			if err := writer.Write(dirPath, iface); err != nil {
				ctx.Cancel(err)
			}
		}(ctx, wg, writer, iface)
	}
	wg.Wait()

	ctx.GeneratorDone()
	return

FuckErr:
	ctx.Cancel(err)
}

// Clone return a copy of Generator
func (g *mirGenerator) Clone() core.Generator {
	return &mirGenerator{
		name:     g.name,
		sinkPath: g.sinkPath,
	}
}

func (g *mirGenerator) cleanup() {
	if g.isCleanup {
		apiPath := path.Join(g.sinkPath, "api")
		core.Logus("cleanup out: %s", apiPath)
		if err := os.RemoveAll(apiPath); err != nil {
			core.Logus("want cleanup out first but failure: %s.do it later by yourself.", err)
		}
	}
}

func generate(generatorName string, sinkPath string, ds core.Descriptors) error {
	var dirPath string

	tmpl, err := templateFrom(generatorName)
	if err != nil {
		return err
	}
	writer := &mirWriter{tmpl: tmpl, ns: naming.NewSnakeNamingStrategy()}
	apiPath := filepath.Join(sinkPath, "api")

FuckErr:
	for key, ifaces := range ds {
		group := ds.GroupFrom(key)
		dirPath = filepath.Join(apiPath, group)
		if err = os.MkdirAll(dirPath, 0755); err != nil {
			break
		}
		filter := utils.NewStrSet()
		hadDecribeCoreInterface := false
		for _, iface := range ifaces.SortedIfaces() {
			var inouts []reflect.Type
			for _, typ := range iface.AllInOuts() {
				if typ.PkgPath() == iface.PkgPath {
					if err := filter.Add(typ.Name()); err == nil {
						inouts = append(inouts, typ)
					}
				} else {
					inouts = append(inouts, typ)
				}
			}
			iface.SetInnerInOuts(inouts)
			if !hadDecribeCoreInterface {
				hadDecribeCoreInterface = true
				iface.SetDeclareCoreInterface(true)
			}
			if err = writer.Write(dirPath, iface); err != nil {
				break FuckErr
			}
		}
	}

	return err
}

func evalSinkPath(path string) (string, error) {
	sp, err := filepath.EvalSymlinks(path)
	if err != nil {
		if os.IsNotExist(err) {
			if !filepath.IsAbs(path) {
				if sp, err = os.Getwd(); err == nil {
					sp = filepath.Join(sp, path)
				}
			} else {
				sp, err = path, nil
			}
		}
	}
	return sp, err
}
