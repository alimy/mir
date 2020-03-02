GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)

.PHONY: build
build: fmt
	go build .

.PHONY: generate
generate:
	-rm -rf mirc/gen
	go generate mirc/main.go

.PHONY: fmt
fmt:
	$(GOFMT) $(GOFILES)
