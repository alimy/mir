// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package generator

import (
	"embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/alimy/mir/v3/core"
	"github.com/alimy/mir/v3/internal/utils"
)

//go:embed templates
var content embed.FS

// tmplInfos generator name map assets name
var tmplInfos = map[string]string{
	core.GeneratorGin:        "templates/gin_iface.tmpl",
	core.GeneratorChi:        "templates/chi_iface.tmpl",
	core.GeneratorMux:        "templates/mux_iface.tmpl",
	core.GeneratorEcho:       "templates/echo_iface.tmpl",
	core.GeneratorIris:       "templates/iris_iface.tmpl",
	core.GeneratorFiber:      "templates/fiber_iface.tmpl",
	core.GeneratorFiberV2:    "templates/fiber_iface_v2.tmpl",
	core.GeneratorMacaron:    "templates/macaron_iface.tmpl",
	core.GeneratorHttpRouter: "templates/httprouter_iface.tmpl",
}

func templateFrom(name string) (*template.Template, error) {
	tmplName, exist := tmplInfos[name]
	if !exist {
		return nil, fmt.Errorf("not exist templates for genererator:%s", name)
	}
	data, err := content.ReadFile(tmplName)
	if err != nil {
		return nil, err
	}
	t := template.New("mir").Funcs(template.FuncMap{
		"notEmptyStr":    notEmptyStr,
		"joinPath":       joinPath,
		"valideQuery":    valideQuery,
		"inflateQuery":   inflateQuery,
		"upperFirstName": upperFirstName,
	})
	if tmpl, err := t.Parse(string(data)); err == nil {
		return tmpl, nil
	} else {
		return nil, err
	}
}

func notEmptyStr(s string) bool {
	return s != ""
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

func upperFirstName(name string) string {
	return utils.UpperFirst(strings.ToLower(name))
}
