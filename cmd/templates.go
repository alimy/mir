// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"embed"
)

//go:embed templates
var fs embed.FS

// tmplCtx template context for generate project
type tmplCtx struct {
	PkgName    string
	MirPkgName string
}

// tmplInfo template data info
type tmplInfo struct {
	name   string
	isTmpl bool
}

func (t tmplInfo) MustBytes() []byte {
	data, err := fs.ReadFile(t.name)
	if err != nil {
		panic(err)
	}
	return data
}

// tmplFiles map of generated file's assets info
var tmplFiles = map[string]map[string]tmplInfo{
	"gin": {
		"Makefile":               {"templates/makefile.tmpl", false},
		"README.md":              {"templates/readme.tmpl", false},
		"go.mod":                 {"templates/gin_go_mod.tmpl", true},
		"main.go":                {"templates/gin_main.tmpl", false},
		"mirc/main.go":           {"templates/gin_mirc_main.tmpl", true},
		"mirc/routes/site.go":    {"templates/gin_mirc_routes_site.tmpl", false},
		"mirc/routes/v1/site.go": {"templates/gin_mirc_routes_site_v1.tmpl", false},
		"mirc/routes/v2/site.go": {"templates/gin_mirc_routes_site_v2.tmpl", false},
	},
	"chi": {
		"Makefile":               {"templates/makefile.tmpl", false},
		"README.md":              {"templates/readme.tmpl", false},
		"go.mod":                 {"templates/chi_go_mod.tmpl", true},
		"main.go":                {"templates/chi_main.tmpl", false},
		"mirc/main.go":           {"templates/chi_mirc_main.tmpl", true},
		"mirc/routes/site.go":    {"templates/chi_mirc_routes_site.tmpl", false},
		"mirc/routes/v1/site.go": {"templates/chi_mirc_routes_site_v1.tmpl", false},
		"mirc/routes/v2/site.go": {"templates/chi_mirc_routes_site_v2.tmpl", false},
	},
	"mux": {
		"Makefile":               {"templates/makefile.tmpl", false},
		"README.md":              {"templates/readme.tmpl", false},
		"go.mod":                 {"templates/mux_go_mod.tmpl", true},
		"main.go":                {"templates/mux_main.tmpl", false},
		"mirc/main.go":           {"templates/mux_mirc_main.tmpl", true},
		"mirc/routes/site.go":    {"templates/mux_mirc_routes_site.tmpl", false},
		"mirc/routes/v1/site.go": {"templates/mux_mirc_routes_site_v1.tmpl", false},
		"mirc/routes/v2/site.go": {"templates/mux_mirc_routes_site_v2.tmpl", false},
	},
	"echo": {
		"Makefile":               {"templates/makefile.tmpl", false},
		"README.md":              {"templates/readme.tmpl", false},
		"go.mod":                 {"templates/echo_go_mod.tmpl", true},
		"main.go":                {"templates/echo_main.tmpl", false},
		"mirc/main.go":           {"templates/echo_mirc_main.tmpl", true},
		"mirc/routes/site.go":    {"templates/echo_mirc_routes_site.tmpl", false},
		"mirc/routes/v1/site.go": {"templates/echo_mirc_routes_site_v1.tmpl", false},
		"mirc/routes/v2/site.go": {"templates/echo_mirc_routes_site_v2.tmpl", false},
	},
	"iris": {
		"Makefile":               {"templates/makefile.tmpl", false},
		"README.md":              {"templates/readme.tmpl", false},
		"go.mod":                 {"templates/iris_go_mod.tmpl", true},
		"main.go":                {"templates/iris_main.tmpl", false},
		"mirc/main.go":           {"templates/iris_mirc_main.tmpl", true},
		"mirc/routes/site.go":    {"templates/iris_mirc_routes_site.tmpl", false},
		"mirc/routes/v1/site.go": {"templates/iris_mirc_routes_site_v1.tmpl", false},
		"mirc/routes/v2/site.go": {"templates/iris_mirc_routes_site_v2.tmpl", false},
	},
	"fiber": {
		"Makefile":               {"templates/makefile.tmpl", false},
		"README.md":              {"templates/readme.tmpl", false},
		"go.mod":                 {"templates/fiber_go_mod.tmpl", true},
		"main.go":                {"templates/fiber_main.tmpl", false},
		"mirc/main.go":           {"templates/fiber_mirc_main.tmpl", true},
		"mirc/routes/site.go":    {"templates/fiber_mirc_routes_site.tmpl", false},
		"mirc/routes/v1/site.go": {"templates/fiber_mirc_routes_site_v1.tmpl", false},
		"mirc/routes/v2/site.go": {"templates/fiber_mirc_routes_site_v2.tmpl", false},
	},
	"fiber-v2": {
		"Makefile":               {"templates/makefile.tmpl", false},
		"README.md":              {"templates/readme.tmpl", false},
		"go.mod":                 {"templates/fiber_v2_go_mod.tmpl", true},
		"main.go":                {"templates/fiber_v2_main.tmpl", false},
		"mirc/main.go":           {"templates/fiber_v2_mirc_main.tmpl", true},
		"mirc/routes/site.go":    {"templates/fiber_mirc_routes_site.tmpl", false},
		"mirc/routes/v1/site.go": {"templates/fiber_mirc_routes_site_v1.tmpl", false},
		"mirc/routes/v2/site.go": {"templates/fiber_mirc_routes_site_v2.tmpl", false},
	},
	"macaron": {
		"Makefile":               {"templates/makefile.tmpl", false},
		"README.md":              {"templates/readme.tmpl", false},
		"go.mod":                 {"templates/macaron_go_mod.tmpl", true},
		"main.go":                {"templates/macaron_main.tmpl", false},
		"mirc/main.go":           {"templates/macaron_mirc_main.tmpl", true},
		"mirc/routes/site.go":    {"templates/macaron_mirc_routes_site.tmpl", false},
		"mirc/routes/v1/site.go": {"templates/macaron_mirc_routes_site_v1.tmpl", false},
		"mirc/routes/v2/site.go": {"templates/macaron_mirc_routes_site_v2.tmpl", false},
	},
	"httprouter": {
		"Makefile":               {"templates/makefile.tmpl", false},
		"README.md":              {"templates/readme.tmpl", false},
		"go.mod":                 {"templates/httprouter_go_mod.tmpl", true},
		"main.go":                {"templates/httprouter_main.tmpl", false},
		"mirc/main.go":           {"templates/httprouter_mirc_main.tmpl", true},
		"mirc/routes/site.go":    {"templates/httprouter_mirc_routes_site.tmpl", false},
		"mirc/routes/v1/site.go": {"templates/httprouter_mirc_routes_site_v1.tmpl", false},
		"mirc/routes/v2/site.go": {"templates/httprouter_mirc_routes_site_v2.tmpl", false},
	},
}
