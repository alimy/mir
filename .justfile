#!/usr/bin/env -S just --justfile

default: help

[doc("list help")]
[group("help")]
help:
  @just --list --justfile {{justfile()}}

[doc("run ci")]
[group("ci")]
ci: vet && test

[doc("test code")]
[group("develop")]
test:
  @echo "Testting code..."
  @go test ./...

[doc("generate documents")]
[group("documents")]
gen-docs:
  @echo "Generating documents..." 
  @-rm -rf docs/public
  @cd docs && hugo --minify  --baseURL "https://alimy.github.io/mir/" && cd -

[doc("run docs")]
[group("documents")]
run-docs:
  @echo "Running docs..."
  @cd docs && hugo serve --minify && cd -

[doc("vet code")]
[group("develop")]
vet:
  @echo "Vetting code..."
  @go vet ./...

[doc("formatting code")]
[group("develop")]
fmt:
  @echo "Formatting code..."
  @go fmt ./...
