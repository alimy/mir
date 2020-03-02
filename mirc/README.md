## Mirc
Mir's help toolkit

### Usage
```bash
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

% cd example
% make generate
% make build
```

### Release
```bash
%  hub release create -m "mirc/{tag eg:v2.0.1} release" mirc/{tag eg:v2.0.1}
```
