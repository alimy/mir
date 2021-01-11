// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

import (
	"embed"
	"text/template"

	"github.com/alimy/embedx"
)

// tmplCtx template context for generate project
type tmplCtx struct {
	PkgName    string
	MirPkgName string
}

// tmplInfo template data info
type tmplInfo struct {
	name string
}

// tmplFiles map of generated file's assets info
var tmplFiles = map[string]map[string]tmplInfo{
	"gin": {
		"Makefile":               {"makefile.tmpl"},
		"README.md":              {"readme.tmpl"},
		"go.mod":                 {"gin_go_mod.tmpl"},
		"main.go":                {"gin_main.tmpl"},
		"mirc/main.go":           {"gin_mirc_main.tmpl"},
		"mirc/routes/site.go":    {"gin_mirc_routes_site.tmpl"},
		"mirc/routes/v1/site.go": {"gin_mirc_routes_site_v1.tmpl"},
		"mirc/routes/v2/site.go": {"gin_mirc_routes_site_v2.tmpl"},
	},
	"chi": {
		"Makefile":               {"makefile.tmpl"},
		"README.md":              {"readme.tmpl"},
		"go.mod":                 {"chi_go_mod.tmpl"},
		"main.go":                {"chi_main.tmpl"},
		"mirc/main.go":           {"chi_mirc_main.tmpl"},
		"mirc/routes/site.go":    {"chi_mirc_routes_site.tmpl"},
		"mirc/routes/v1/site.go": {"chi_mirc_routes_site_v1.tmpl"},
		"mirc/routes/v2/site.go": {"chi_mirc_routes_site_v2.tmpl"},
	},
	"mux": {
		"Makefile":               {"makefile.tmpl"},
		"README.md":              {"readme.tmpl"},
		"go.mod":                 {"mux_go_mod.tmpl"},
		"main.go":                {"mux_main.tmpl"},
		"mirc/main.go":           {"mux_mirc_main.tmpl"},
		"mirc/routes/site.go":    {"mux_mirc_routes_site.tmpl"},
		"mirc/routes/v1/site.go": {"mux_mirc_routes_site_v1.tmpl"},
		"mirc/routes/v2/site.go": {"mux_mirc_routes_site_v2.tmpl"},
	},
	"echo": {
		"Makefile":               {"makefile.tmpl"},
		"README.md":              {"readme.tmpl"},
		"go.mod":                 {"echo_go_mod.tmpl"},
		"main.go":                {"echo_main.tmpl"},
		"mirc/main.go":           {"echo_mirc_main.tmpl"},
		"mirc/routes/site.go":    {"echo_mirc_routes_site.tmpl"},
		"mirc/routes/v1/site.go": {"echo_mirc_routes_site_v1.tmpl"},
		"mirc/routes/v2/site.go": {"echo_mirc_routes_site_v2.tmpl"},
	},
	"iris": {
		"Makefile":               {"makefile.tmpl"},
		"README.md":              {"readme.tmpl"},
		"go.mod":                 {"iris_go_mod.tmpl"},
		"main.go":                {"iris_main.tmpl"},
		"mirc/main.go":           {"iris_mirc_main.tmpl"},
		"mirc/routes/site.go":    {"iris_mirc_routes_site.tmpl"},
		"mirc/routes/v1/site.go": {"iris_mirc_routes_site_v1.tmpl"},
		"mirc/routes/v2/site.go": {"iris_mirc_routes_site_v2.tmpl"},
	},
	"fiber": {
		"Makefile":               {"makefile.tmpl"},
		"README.md":              {"readme.tmpl"},
		"go.mod":                 {"fiber_go_mod.tmpl"},
		"main.go":                {"fiber_main.tmpl"},
		"mirc/main.go":           {"fiber_mirc_main.tmpl"},
		"mirc/routes/site.go":    {"fiber_mirc_routes_site.tmpl"},
		"mirc/routes/v1/site.go": {"fiber_mirc_routes_site_v1.tmpl"},
		"mirc/routes/v2/site.go": {"fiber_mirc_routes_site_v2.tmpl"},
	},
	"fiber-v2": {
		"Makefile":               {"makefile.tmpl"},
		"README.md":              {"readme.tmpl"},
		"go.mod":                 {"fiber_v2_go_mod.tmpl"},
		"main.go":                {"fiber_v2_main.tmpl"},
		"mirc/main.go":           {"fiber_v2_mirc_main.tmpl"},
		"mirc/routes/site.go":    {"fiber_mirc_routes_site.tmpl"},
		"mirc/routes/v1/site.go": {"fiber_mirc_routes_site_v1.tmpl"},
		"mirc/routes/v2/site.go": {"fiber_mirc_routes_site_v2.tmpl"},
	},
	"macaron": {
		"Makefile":               {"makefile.tmpl"},
		"README.md":              {"readme.tmpl"},
		"go.mod":                 {"macaron_go_mod.tmpl"},
		"main.go":                {"macaron_main.tmpl"},
		"mirc/main.go":           {"macaron_mirc_main.tmpl"},
		"mirc/routes/site.go":    {"macaron_mirc_routes_site.tmpl"},
		"mirc/routes/v1/site.go": {"macaron_mirc_routes_site_v1.tmpl"},
		"mirc/routes/v2/site.go": {"macaron_mirc_routes_site_v2.tmpl"},
	},
	"httprouter": {
		"Makefile":               {"makefile.tmpl"},
		"README.md":              {"readme.tmpl"},
		"go.mod":                 {"httprouter_go_mod.tmpl"},
		"main.go":                {"httprouter_main.tmpl"},
		"mirc/main.go":           {"httprouter_mirc_main.tmpl"},
		"mirc/routes/site.go":    {"httprouter_mirc_routes_site.tmpl"},
		"mirc/routes/v1/site.go": {"httprouter_mirc_routes_site_v1.tmpl"},
		"mirc/routes/v2/site.go": {"httprouter_mirc_routes_site_v2.tmpl"},
	},
}

func newTemplate() (*template.Template, error) {
	//go:embed templates
	var content embed.FS

	embedFS := embedx.ChangeRoot(content, "templates")
	t := template.New("mirc").Funcs(template.FuncMap{
		"notEmptyStr": notEmptyStr,
	})
	return embedx.ParseWith(t, embedFS, "*.tmpl")
}

func notEmptyStr(s string) bool {
	return s != ""
}
