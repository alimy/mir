#!/usr/bin/env -S just --justfile

TARGET := if os_family() == "windows" { "mir-exampe.exe" } else { "mir-example" }

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
  @go build -o {{TARGET}}

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

[doc("clean project")]
[group("develop")]
clean:
  @echo "Clean project..."
  @-rm -f {{TARGET}}
