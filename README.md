# Mir
[![Build Status](https://api.travis-ci.com/alimy/mir.svg?branch=master)](https://travis-ci.com/alimy/mir)
[![codecov](https://codecov.io/gh/alimy/mir/branch/master/graph/badge.svg)](https://codecov.io/gh/alimy/mir)
[![GoDoc](https://godoc.org/github.com/alimy/mir?status.svg)](https://godoc.org/github.com/alimy/mir)

Mir is used for register handler to http router(eg: [Gin](https://github.com/gin-gonic/gin), [Chi](https://github.com/go-chi/chi), [Echo](https://github.com/labstack/echo), [Iris](https://github.com/kataras/iris), [Macaron](https://github.com/go-macaron/macaron), [Mux](https://github.com/gorilla/mux), [httprouter](https://github.com/julienschmidt/httprouter))
 depends on struct tag string info that defined in logic object's struct type field.
 
 ## Usage
 ```
% go get github.com/alimy/mir/mirc/v2@latest
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
% ./mir-example
 
 ```
