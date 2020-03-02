## Mirc
Mir's help toolkit

### Usage
```bash
% mir help
mir help tookit

Usage:
  mir [command]

Available Commands:
  help        Help about any command
  new         create template project
  version     show version information

Flags:
  -h, --help   help for mir

Use "mir [command] --help" for more information about a command.

% mir help new
create template project

Usage:
  mir new [flags]

Flags:
  -d, --dst string    genereted destination target directory (default ".")
  -h, --help          help for new
  -t, --type string   generated engine type style(eg: gin,chi,mux,httprout) (default "gin")

% mir new -d example
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

4 directories, 8 files

% cd example
% make generate
% make build
```

### Release
```bash
%  hub release create -m "mirc/{tag eg:v2.0.1} release" mirc/{tag eg:v2.0.1}
```
