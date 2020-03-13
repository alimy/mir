package generator

import (
	"unsafe"

	"github.com/alimy/mir/v2/core"
)

//go:generate go-bindata -nomemcopy -pkg=${GOPACKAGE} -ignore=README.md -prefix=templates -debug=false -o=templates_gen.go templates/...

// tmplFiles generator name map assets name
var tmplFiles = map[string]string{
	core.GeneratorGin:        "gin_iface.tmpl",
	core.GeneratorChi:        "chi_iface.tmpl",
	core.GeneratorMux:        "mux_iface.tmpl",
	core.GeneratorEcho:       "echo_iface.tmpl",
	core.GeneratorMacaron:    "macaron_iface.tmpl",
	core.GeneratorHttpRouter: "httprouter_iface.tmpl",
}

// bytesToString convert byte slice to string
func bytesToString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}
