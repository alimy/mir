// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package cmd

//go:generate go-bindata -nomemcopy -pkg=${GOPACKAGE} -ignore=README.md -prefix=templates -debug=false -o=templates_gen.go templates/...

// tmplCtx template context for generate project
type tmplCtx struct {
	PkgName string
}

// tmplInfo template data info
type tmplInfo struct {
	name   string
	isTmpl bool
}

// tmplFiles map of generated file's assets info
var tmplFiles = map[string]map[string]tmplInfo{
	"gin": {
		"Makefile":               {"makefile.tmpl", false},
		"README.md":              {"readme.tmpl", false},
		"go.mod":                 {"gin_go_mod.tmpl", true},
		"main.go":                {"gin_main.tmpl", false},
		"mirc/main.go":           {"gin_mirc_main.tmpl", true},
		"mirc/routes/site.go":    {"gin_mirc_routes_site.tmpl", false},
		"mirc/routes/v1/site.go": {"gin_mirc_routes_site_v1.tmpl", false},
		"mirc/routes/v2/site.go": {"gin_mirc_routes_site_v2.tmpl", false},
	},
	"chi": {
		"Makefile":               {"makefile.tmpl", false},
		"README.md":              {"readme.tmpl", false},
		"go.mod":                 {"chi_go_mod.tmpl", true},
		"main.go":                {"chi_main.tmpl", false},
		"mirc/main.go":           {"chi_mirc_main.tmpl", true},
		"mirc/routes/site.go":    {"chi_mirc_routes_site.tmpl", false},
		"mirc/routes/v1/site.go": {"chi_mirc_routes_site_v1.tmpl", false},
		"mirc/routes/v2/site.go": {"chi_mirc_routes_site_v2.tmpl", false},
	},
	"mux": {
		"Makefile":               {"makefile.tmpl", false},
		"README.md":              {"readme.tmpl", false},
		"go.mod":                 {"mux_go_mod.tmpl", true},
		"main.go":                {"mux_main.tmpl", false},
		"mirc/main.go":           {"mux_mirc_main.tmpl", true},
		"mirc/routes/site.go":    {"mux_mirc_routes_site.tmpl", false},
		"mirc/routes/v1/site.go": {"mux_mirc_routes_site_v1.tmpl", false},
		"mirc/routes/v2/site.go": {"mux_mirc_routes_site_v2.tmpl", false},
	},
	"echo": {
		"Makefile":               {"makefile.tmpl", false},
		"README.md":              {"readme.tmpl", false},
		"go.mod":                 {"echo_go_mod.tmpl", true},
		"main.go":                {"echo_main.tmpl", false},
		"mirc/main.go":           {"echo_mirc_main.tmpl", true},
		"mirc/routes/site.go":    {"echo_mirc_routes_site.tmpl", false},
		"mirc/routes/v1/site.go": {"echo_mirc_routes_site_v1.tmpl", false},
		"mirc/routes/v2/site.go": {"echo_mirc_routes_site_v2.tmpl", false},
	},
	"iris": {
		"Makefile":               {"makefile.tmpl", false},
		"README.md":              {"readme.tmpl", false},
		"go.mod":                 {"iris_go_mod.tmpl", true},
		"main.go":                {"iris_main.tmpl", false},
		"mirc/main.go":           {"iris_mirc_main.tmpl", true},
		"mirc/routes/site.go":    {"iris_mirc_routes_site.tmpl", false},
		"mirc/routes/v1/site.go": {"iris_mirc_routes_site_v1.tmpl", false},
		"mirc/routes/v2/site.go": {"iris_mirc_routes_site_v2.tmpl", false},
	},
	"macaron": {
		"Makefile":               {"makefile.tmpl", false},
		"README.md":              {"readme.tmpl", false},
		"go.mod":                 {"macaron_go_mod.tmpl", true},
		"main.go":                {"macaron_main.tmpl", false},
		"mirc/main.go":           {"macaron_mirc_main.tmpl", true},
		"mirc/routes/site.go":    {"macaron_mirc_routes_site.tmpl", false},
		"mirc/routes/v1/site.go": {"macaron_mirc_routes_site_v1.tmpl", false},
		"mirc/routes/v2/site.go": {"macaron_mirc_routes_site_v2.tmpl", false},
	},
	"httprouter": {
		"Makefile":               {"makefile.tmpl", false},
		"README.md":              {"readme.tmpl", false},
		"go.mod":                 {"httprouter_go_mod.tmpl", true},
		"main.go":                {"httprouter_main.tmpl", false},
		"mirc/main.go":           {"httprouter_mirc_main.tmpl", true},
		"mirc/routes/site.go":    {"httprouter_mirc_routes_site.tmpl", false},
		"mirc/routes/v1/site.go": {"httprouter_mirc_routes_site_v1.tmpl", false},
		"mirc/routes/v2/site.go": {"httprouter_mirc_routes_site_v2.tmpl", false},
	},
}
