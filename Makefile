GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)

.PHONY: default
default: ci

.PHONY: ci
ci: misspell vet
	go test ./...

.PHONY: build
build: fmt
	go build -o mir mirc/main.go

.PHONY: test
test:
	go test ./...

.PHONY: gen-docs
gen-docs: 
	@-rm -rf docs/public
	@cd docs && hugo --minify  --baseURL "https://alimy.github.io/mir/" && cd -

.PHONY: run-docs
run-docs: 
	@cd docs && hugo serve --minify && cd -

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt-check
fmt-check:
	@diff=$$($(GOFMT) -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

.PHONY: fmt
fmt:
	$(GOFMT) $(GOFILES)
