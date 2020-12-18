// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package generator

import (
	"embed"
	"unsafe"

	"github.com/alimy/mir/v2/core"
)

//go:embed templates
var fs embed.FS

type tmplInfos map[string]string

// tmplFiles generator name map assets name
var tmplFiles = tmplInfos{
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

func (t tmplInfos) notExist(name string) bool {
	if _, exist := t[name]; !exist {
		return true
	}
	return false
}

func (t tmplInfos) mustString(name string) string {
	data, err := fs.ReadFile(t[name])
	if err != nil {
		panic(err)
	}
	return *(*string)(unsafe.Pointer(&data))
}
