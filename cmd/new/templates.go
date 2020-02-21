package new

import "github.com/alimy/mir/v2/core"

//go:generate go-bindata -nomemcopy -pkg=${GOPACKAGE} -ignore=README.md -prefix=templates -debug=false -o=templates_gen.go templates/...

var tmplFiles = map[string]map[string]string{
	core.EngineGin: {
		"Makefile":               "makefile.tmpl",
		"README.md":              "readme.tmpl",
		"go.mod":                 "gin_go_mod.tmpl",
		"main.go":                "gin_main.tmpl",
		"mirc/main.go":           "gin_mirc_main.tmpl",
		"mirc/routes/site.go":    "gin_mirc_routes_site.tmpl",
		"mirc/routes/v1/site.go": "gin_mirc_routes_site_v1.tmpl",
		"mirc/routes/v2/site.go": "gin_mirc_routes_site_v2.tmpl",
	},
}
