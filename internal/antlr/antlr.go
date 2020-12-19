// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.

package antlr

//go:generate java -jar ./jar/antlr-4.9-complete.jar -o parser -visitor -Dlanguage=Go Mir.g4
//go:generate sh -c "(cd parser; sed -i '' -e 's!github.com/antlr/antlr4/runtime/Go/antlr!github.com/alimy/antlr4-go!' *.go )"
