## Mirc
Mir's help toolkit

### Usage
```bash
% go get github.com/alimy/mir/mirc/v2@latest
% mirc new -h
create template project

Usage:
  mirc new [flags]

Flags:
  -d, --dst string     genereted destination target directory (default ".")
  -h, --help           help for new
      --mir string     mir replace package name or place
  -p, --pkg string     project's package name (default "github.com/alimy/mir-example")
  -s, --style string   generated engine style eg: gin,chi,mux,echo,iris,fiber,macaron,httprouter (default "gin")

% mirc new -d example 
% tree example
example
├── Makefile
├── README.md
├── go.mod
├── main.go
└── mirc
    ├── main.go
    └── routes
        ├── site.go
        ├── v1
        │   └── site.go
        └── v2
            └── site.go

% cd example
% make generate
% make build
```
