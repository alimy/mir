package generator

//go:generate go-bindata -nomemcopy -pkg=${GOPACKAGE} -ignore=README.md -prefix=templates -debug=false -o=templates_gen.go templates/...
