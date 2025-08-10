#!/usr/bin/env -S just --justfile

MOD_NAME := "github.com/alimy/mir/sail-example/v5"
TARGET := if os_family() == "windows" { "sail-example.exe" } else { "sail-example" }
TAGS := "go_json"
BUILD_VERSION := `git describe --tags --always`
BUILD_DATE := `date +'%Y-%m-%d %H:%M:%S %Z'`
SHA_SHORT := `git rev-parse --short HEAD`

LDFLAGS := shell('echo \
  "-X $1$2/internal/conf.version=$3$1 \
   -X $1$2/internal/conf.buildDate=$4$1 \
   -X $1$2/internal/conf.commitID=$5$1 \
   -X $1$2/internal/conf.buildTags=$6$1"', 
  "'", 
  MOD_NAME, 
  BUILD_VERSION, 
  BUILD_DATE,
  SHA_SHORT,
  TAGS,
)

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
  @go build -trimpath -tags '{{TAGS}}' -o {{TARGET}} -ldflags "{{LDFLAGS}}"

[doc("run code")]
[group("develop")]
run:
  @echo "Running code..."
  @go run -trimpath -tags '{{TAGS}}' -ldflags "{{LDFLAGS}}" main.go serve

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
  go fmt ./auto/...

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
