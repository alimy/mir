#!/usr/bin/env -S just --justfile

TARGET := if os_family() == "windows" { "sail-example.exe" } else { "sail-example" }
TAGS := "go_json"
LDFLAGS := ```
  MOD_NAME="github.com/alimy/mir/sail-example/v5"
  TAGS="go_json"
  BUILD_VERSION=$(git describe --tags --always)
  BUILD_DATE=$(date +'%Y-%m-%d %H:%M:%S %Z')
  SHA_SHORT=$(git rev-parse --short HEAD)
  echo "-X '$MOD_NAME/internal/conf.version=$BUILD_VERSION' 
        -X '$MOD_NAME/internal/conf.buildDate=$BUILD_DATE' 
        -X '$MOD_NAME/internal/conf.commitID=$SHA_SHORT'  
        -X '$MOD_NAME/internal/conf.buildTags=$TAGS'"
```

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
