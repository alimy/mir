GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)
TAGS ?= go_json
LDFLAGS = -w -s

.PHONY: default
default: run

.PHONY: build
build: fmt
	go build -trimpath -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o mir-examples main.go

.PHONY: build
run: fmt
	go run -trimpath -tags '$(TAGS)' -ldflags '$(LDFLAGS)' main.go

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
