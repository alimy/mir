#!/usr/bin/env -S just --justfile

target := if os_family() == "windows" { "exampes.exe" } else { "examples" }

alias serve := run
alias gen := generate
alias gen-mir := generate

default: help

[doc("list help")]
[group("help")]
help:
  @just --list --justfile {{justfile()}}

[doc("build code")]
[group("develop")]
build:
  @echo "Building code..."
  @go build -o {{target}}

[doc("run code")]
[group("develop")]
run:
  @echo "Running code..."
  @go run .

[doc("tidy mod")]
[group("develop")]
mod-tidy:
  @echo "Tidy mod"
  @go mod tidy

[doc("generate code")]
[group("develop")]
generate:
  @echo "Generatting code..."
  @go generate mirc/gen.go
  go fmt ./mirc/auto/...

[doc("test code")]
[group("develop")]
test:
  @echo "Testting code..."
  @go test ./...

[doc("vet code")]
[group("develop")]
vet:
  @echo "Vetting code..."
  @go vet ./...

[doc("format code")]
[group("develop")]
fmt:
  @echo "Formatting code..."
  @go fmt ./...
