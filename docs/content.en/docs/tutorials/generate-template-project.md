---
weight: 20
title: "Generate Template Project"
---

## Generate Template Project
```bash
% go install github.com/alimy/mir/mirc/v3@latest
% mirc new -h
create template project

Usage:
  mirc new [flags]

Flags:
  -d, --dst string     genereted destination target directory (default ".")
  -h, --help           help for new
      --mir string     mir replace package name or place
  -p, --pkg string     project's package name (default "github.com/alimy/mir-example")
  -s, --style string   generated engine style eg: gin,chi,mux,hertz,echo,iris,fiber,fiber-v2,macaron,httprouter (default "gin")

% mirc new -d example 
% tree example
example
.
|-- Makefile
|-- README.md
|-- go.mod
|-- go.sum
|-- main.go
|-- mirc
|   |-- auto
|   |   `-- api
|   |       |-- site.go
|   |       |-- v1
|   |       |   `-- site.go
|   |       `-- v2
|   |           `-- site.go
|   |-- gen.go
|   `-- routes
|       |-- site.go
|       |-- v1
|       |   `-- site.go
|       `-- v2
|           `-- site.go
`-- servants
    |-- core.go
    |-- servants.go
    |-- site.go
    |-- site_v1.go
    `-- site_v2.go

% cd example
% make generate
% make run
```
