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
}
