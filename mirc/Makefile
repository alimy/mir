GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)

LDFLAGS += -X "github.com/alimy/mir/mirc/v5/version.GitHash=$(shell git rev-parse --short=12 HEAD)"

RELEASE_ROOT = release
RELEASE_LINUX_AMD64 = $(RELEASE_ROOT)/linux-amd64/mirc
RELEASE_DARWIN_AMD64 = $(RELEASE_ROOT)/darwin-amd64/mirc
RELEASE_DARWIN_ARM64 = $(RELEASE_ROOT)/darwin-arm64/mirc
RELEASE_WINDOWS_AMD64 = $(RELEASE_ROOT)/windows-amd64/mirc

.PHONY: build
build: fmt
	go build  -ldflags '$(LDFLAGS)' -o mirc main.go

.PHONY: release
release: linux-amd64 darwin-amd64 darwin-arm64 windows-x64
	cp ../LICENSE README.md $(RELEASE_LINUX_AMD64)
	cp ../LICENSE README.md $(RELEASE_DARWIN_AMD64)
	cp ../LICENSE README.md $(RELEASE_DARWIN_ARM64)
	cp ../LICENSE README.md $(RELEASE_WINDOWS_AMD64)
	cd $(RELEASE_LINUX_AMD64)/.. && rm -f *.zip && zip -r mirc-linux_amd64.zip mirc && cd -
	cd $(RELEASE_DARWIN_AMD64)/.. && rm -f *.zip && zip -r mirc-darwin_amd64.zip mirc && cd -
	cd $(RELEASE_DARWIN_ARM64)/.. && rm -f *.zip && zip -r mirc-darwin_arm64.zip mirc && cd -
	cd $(RELEASE_WINDOWS_AMD64)/.. && rm -f *.zip && zip -r mirc-windows_amd64.zip mirc && cd -

.PHONY: linux-amd64
linux-amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -ldflags '$(LDFLAGS)' -o $(RELEASE_LINUX_AMD64)/mirc main.go

.PHONY: darwin-amd64
darwin-amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build  -ldflags '$(LDFLAGS)' -o $(RELEASE_DARWIN_AMD64)/mirc main.go

.PHONY: darwin-arm64
darwin-arm64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build  -ldflags '$(LDFLAGS)' -o $(RELEASE_DARWIN_ARM64)/mirc main.go

.PHONY: windows-x64
windows-x64:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -ldflags '$(LDFLAGS)' -o $(RELEASE_WINDOWS_AMD64)/mirc main.go

.PHONY: fmt
fmt:
	$(GOFMT) $(GOFILES)

.PHONY: check-all
check-all: check-debug-all

STYLES = chi hertz echo fiber gin httprouter iris macaron mux
.PHONY: check-debug-all
check-debug-all:
	@for target in $(STYLES); do \
	  echo "==========[ processing $$target ]=========="; \
	  rm -rf $$target; \
	  ./mirc new --mir ../../ --style $$target --pkg github.com/alimy/mir-$$target --dst $$target; \
	  cd $$target; \
	  make mod-tidy; \
	  cd -; \
	  echo ""; \
	done

.PHONY: check-release-all
check-release-all:
	@for target in $(STYLES); do \
	  echo "==========[ processing $$target ]=========="; \
	  rm -rf $$target; \
	  ./mirc new --style $$target --pkg github.com/alimy/mir-$$target --dst $$target; \
	  cd $$target; \
	  make mod-tidy; \
	  cd -; \
	  echo ""; \
	done

.PHONY: clean
clean:
	-rm -rf $(STYLES)
