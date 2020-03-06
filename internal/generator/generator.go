// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package generator

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/alimy/mir/v2"
	"github.com/alimy/mir/v2/core"
)

var (
	errNotExistSinkPath = errors.New("not exist output path")
)

func init() {
	core.RegisterGenerators(
		mirGenerator{name: core.GeneratorGin},
		mirGenerator{name: core.GeneratorChi},
		mirGenerator{name: core.GeneratorMux},
		mirGenerator{name: core.GeneratorHttpRouter},
	)
}

type mirGenerator struct {
	sinkPath string
	name     string
}

// Name name of generator
func (g mirGenerator) Name() string {
	return g.name
}

// Init init generator
func (g mirGenerator) Init(opts core.InitOpts) (err error) {
	if len(opts) != 0 {
		if sp, exist := opts[core.OptSinkPath]; exist {
			g.sinkPath, err = evalSinkPath(sp)
			return
		}
	}
	return errNotExistSinkPath
}

// Generate generate interface code
func (g mirGenerator) Generate(ds core.Descriptors) error {
	return generate(g.name, g.sinkPath, ds)
}

func notEmptyStr(s string) bool {
	return s != ""
}

func notHttpAny(m string) bool {
	return m != mir.MethodAny
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

func generate(generatorName string, sinkPath string, ds core.Descriptors) error {
	var (
		err               error
		dirPath, filePath string
	)

	apiPath := filepath.Join(sinkPath, "api")
	tmpl := template.New("mir").Funcs(template.FuncMap{
		"notEmptyStr":  notEmptyStr,
		"notHttpAny":   notHttpAny,
		"join":         path.Join,
		"valideQuery":  valideQuery,
		"inflateQuery": inflateQuery,
	})
	assetName, exist := tmplFiles[generatorName]
	if !exist {
		return fmt.Errorf("not exist templates for genererator:%s", generatorName)
	}
	if tmpl, err = tmpl.Parse(MustAssetString(assetName)); err != nil {
		return err
	}

FuckErr:
	for key, ifaces := range ds {
		group := ds.GroupFrom(key)
		dirPath = filepath.Join(apiPath, ds.SnakeStr(group))
		if err = os.MkdirAll(dirPath, 0755); err != nil {
			break
		}
		for _, iface := range ifaces {
			filePath = filepath.Join(dirPath, iface.SnakeFileName())
			file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				break FuckErr
			}
			if err = tmpl.Execute(file, iface); err != nil {
				break FuckErr
			}
			if err = file.Close(); err != nil {
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
