// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package generator

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"

	"github.com/alimy/mir/v2"
	"github.com/alimy/mir/v2/core"
	"github.com/alimy/mir/v2/internal/container"
)

var (
	errNotExistSinkPath = errors.New("not exist output path")
)

func init() {
	core.RegisterGenerators(
		&mirGenerator{name: core.GeneratorGin},
		&mirGenerator{name: core.GeneratorChi},
		&mirGenerator{name: core.GeneratorMux},
		&mirGenerator{name: core.GeneratorMacaron},
		&mirGenerator{name: core.GeneratorHttpRouter},
	)
}

type mirGenerator struct {
	sinkPath string
	name     string
}

// Name name of generator
func (g *mirGenerator) Name() string {
	return g.name
}

// Init init generator
func (g *mirGenerator) Init(opts core.InitOpts) (err error) {
	if len(opts) != 0 {
		if sp, exist := opts[core.OptSinkPath]; exist {
			g.sinkPath, err = evalSinkPath(sp)
			return
		}
	}
	return errNotExistSinkPath
}

// Generate serial generate interface code
func (g *mirGenerator) Generate(ds core.Descriptors) error {
	return generate(g.name, g.sinkPath, ds)
}

// GoGenerator concurrent generate interface code
func (g *mirGenerator) GoGenerate(ctx core.MirCtx) {
	tmpl, err := templateFrom(g.name)
	if err != nil {
		ctx.Cancel(err)
		return
	}
	apiPath := filepath.Join(g.sinkPath, "api")
	ifaceSource, _ := ctx.Pipe()
	onceSet := container.NewOnceSet(func(path string) error {
		return os.MkdirAll(path, 0755)
	})

	wg := &sync.WaitGroup{}
	for iface := range ifaceSource {
		select {
		case <-ctx.Done():
			return
		default:
			go goGenerate(ctx, wg, tmpl, onceSet, apiPath, iface)
		}
	}
	wg.Wait()

	ctx.GeneratorDone()
}

// Clone return a copy of Generator
func (g *mirGenerator) Clone() core.Generator {
	return &mirGenerator{
		name:     g.name,
		sinkPath: g.sinkPath,
	}
}

func notEmptyStr(s string) bool {
	return s != ""
}

func notHttpAny(m string) bool {
	return m != mir.MethodAny
}

func joinPath(group, subpath string) string {
	if group == "" {
		return subpath
	}
	b := &strings.Builder{}
	if !strings.HasPrefix(group, "/") {
		b.WriteByte('/')
	}
	b.WriteString(group)
	if !strings.HasSuffix(group, "/") && !strings.HasPrefix(subpath, "/") {
		b.WriteByte('/')
	}
	b.WriteString(subpath)
	return b.String()
}

func valideQuery(qs []string) bool {
	size := len(qs)
	return size != 0 && size%2 == 0
}

func inflateQuery(qs []string) string {
	var b strings.Builder
	last := len(qs) - 1
	b.Grow(last * 10)
	for _, s := range qs {
		b.WriteRune('"')
		b.WriteString(s)
		b.WriteString(`",`)
	}
	return strings.TrimRight(b.String(), ",")
}

func templateFrom(generatorName string) (*template.Template, error) {
	tmpl := template.New("mir").Funcs(template.FuncMap{
		"notEmptyStr":  notEmptyStr,
		"notHttpAny":   notHttpAny,
		"joinPath":     joinPath,
		"valideQuery":  valideQuery,
		"inflateQuery": inflateQuery,
	})
	assetName, exist := tmplFiles[generatorName]
	if !exist {
		return nil, fmt.Errorf("not exist templates for genererator:%s", generatorName)
	}
	return tmpl.Parse(MustAssetString(assetName))
}

func generate(generatorName string, sinkPath string, ds core.Descriptors) error {
	var dirPath string

	tmpl, err := templateFrom(generatorName)
	if err != nil {
		return err
	}
	apiPath := filepath.Join(sinkPath, "api")

FuckErr:
	for key, ifaces := range ds {
		group := ds.GroupFrom(key)
		dirPath = filepath.Join(apiPath, ds.SnakeStr(group))
		if err = os.MkdirAll(dirPath, 0755); err != nil {
			break
		}
		for _, iface := range ifaces {
			if err = doGenerate(dirPath, tmpl, iface); err != nil {
				break FuckErr
			}
		}
	}

	return err
}

func goGenerate(ctx core.MirCtx, wg *sync.WaitGroup, tmpl *template.Template,
	onceSet container.OnceSet, apiPath string, iface *core.IfaceDescriptor) {
	defer wg.Done()

	dirPath := filepath.Join(apiPath, iface.SnakeGroup())
	err := onceSet.Add(dirPath)
	if err == nil {
		err = doGenerate(dirPath, tmpl, iface)
	}
	if err != nil {
		ctx.Cancel(err)
	}
}

func doGenerate(dirPath string, tmpl *template.Template, iface *core.IfaceDescriptor) error {
	filePath := filepath.Join(dirPath, iface.SnakeFileName())
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err == nil {
		defer func() {
			_ = file.Close()
		}()
		err = tmpl.Execute(file, iface)
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
