GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)
TAGS ?= go_json

BUILD_VERSION := $(shell git describe --tags --always)
BUILD_DATE := $(shell date +'%Y-%m-%d %H:%M:%S %Z')
SHA_SHORT := $(shell git rev-parse --short HEAD)

MOD_NAME = github.com/alimy/mir/sail-example/v4
LDFLAGS = -X "${MOD_NAME}/internal/conf.version=${BUILD_VERSION}" \
          -X "${MOD_NAME}/internal/conf.buildDate=${BUILD_DATE}" \
          -X "${MOD_NAME}/internal/conf.commitID=${SHA_SHORT}" \
          -X "${MOD_NAME}/internal/conf.buildTags=${TAGS}" \
		  -w -s

.PHONY: default
default: serve

.PHONY: build
build: fmt
	go build -trimpath -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o sail-example main.go

.PHONY: build
serve: fmt
	go run -trimpath -tags '$(TAGS)' -ldflags '$(LDFLAGS)' main.go serve

.PHONY: mod-tidy
mod-tidy:
	@go mod download
	@go mod tidy

.PHONY: generate
generate:
	@go generate mirc/gen.go
	@$(GOFMT) ./

.PHONY: fmt
fmt:
	$(GOFMT) $(GOFILES)
