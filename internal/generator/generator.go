// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package generator

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/alimy/mir/v2"
	"github.com/alimy/mir/v2/core"
)

func init() {
	core.RegisterGenerators(generatorGin{},
		generatorChi{},
		generatorMux{},
		generatorHttpRouter{})
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

func generate(ds core.Descriptors, opts *core.Options) error {
	var (
		err               error
		dirPath, filePath string
	)

	apiPath := filepath.Join(opts.SinkPath(), "api")
	tmpl := template.New("mir").Funcs(template.FuncMap{
		"notEmptyStr":  notEmptyStr,
		"notHttpAny":   notHttpAny,
		"join":         path.Join,
		"valideQuery":  valideQuery,
		"inflateQuery": inflateQuery,
	})
	assetName, exist := tmplFiles[opts.GeneratorName]
	if !exist {
		return fmt.Errorf("not exist templates for genererator:%s", opts.GeneratorName)
	}
	if tmpl, err = tmpl.Parse(MustAssetString(assetName)); err != nil {
		return err
	}

FuckErr:
	for key, ifaces := range ds {
		group := ds.GroupFrom(key)
		dirPath = filepath.Join(apiPath, group)
		if err = os.MkdirAll(dirPath, 0755); err != nil {
			break
		}
		for _, iface := range ifaces {
			filePath = filepath.Join(apiPath, iface.SnakeFileName())
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
