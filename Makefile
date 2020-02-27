GO ?= go
GOFMT ?= gofmt -s
PACKAGES ?= $(shell GO111MODULE=on $(GO) list ./...)
GOFILES := $(shell find . -name "*.go" -type f)

.PHONY: default
default: ci

.PHONY: ci
ci: misspell vet test

.PHONY: build
build: fmt
	go build -o mir cmd/mir.go

.PHONY: generate
generate:
	-rm -f generator/templates_gen.go
	-rm -f cmd/create/templates_gen.go
	go generate generator/templates.go
	go generate cmd/create/templates.go
	$(GOFMT) -w generator/templates.go
	$(GOFMT) -w cmd/create/templates.go

.PHONY: test
test: fmt
	go test .

.PHONY: fmt
fmt:
	$(GOFMT) -w $(GOFILES)

.PHONY: fmt-check
fmt-check:
	@diff=$$($(GOFMT) -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

.PHONY: vet
vet:
	go vet .

.PHONY: lint
lint:
	@hash golint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u golang.org/x/lint/golint; \
	fi
	for PKG in $(PACKAGES); do golint -min_confidence 1.0 -set_exit_status $$PKG || exit 1; done;

.PHONY: misspell-check
misspell-check:
	@hash misspell > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		GO111MODULE=off $(GO) get -u github.com/client9/misspell/cmd/misspell; \
	fi
	misspell -error $(GOFILES)

.PHONY: misspell
misspell:
	@hash misspell > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		GO111MODULE=off $(GO) get -u github.com/client9/misspell/cmd/misspell; \
	fi
	misspell -w $(GOFILES)

.PHONY: tools
tools:
	GO111MODULE=off $(GO) get golang.org/x/lint/golint
	GO111MODULE=off $(GO) get github.com/client9/misspell/cmd/misspell
	GO111MODULE=off $(GO) get github.com/onsi/ginkgo/ginkgo
