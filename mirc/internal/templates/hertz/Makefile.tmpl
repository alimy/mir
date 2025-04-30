GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)

.PHONY: default
default: run

.PHONY: build
build: fmt
	go build -o mir-examples main.go

.PHONY: build
run: fmt
	go run main.go

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
