// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package mux

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/alimy/mir"
	"github.com/gorilla/mux"
)

var _ mir.Engine = &mirEngine{}

// mirEngine used to implements mir.Engine interface
type mirEngine struct {
	engine *mux.Router
}

// Register register entries to gin engine
func (e *mirEngine) Register(entries []*mir.TagMir) error {
	for _, entry := range entries {
		var router *mux.Router
		if entry.Group == "" || entry.Group == "/" {
			router = e.engine
		} else {
			pathPrefix := entry.Group
			if !strings.HasPrefix(entry.Group, "/") {
				pathPrefix = "/" + entry.Group
			}
			router = e.engine.PathPrefix(pathPrefix).Subrouter()
		}
		if err := handlerChainTo(router, entry.Chain); err != nil {
			return err
		}
		// Notice just return if catch a error or continue next entry register
		if err := registerWith(router, entry.Fields); err != nil {
			return err
		}
	}
	return nil
}

// registerWith register fields to give router
func registerWith(router *mux.Router, fields []*mir.TagField) error {
	for _, field := range fields {
		if handlerFunc, ok := field.Handler.(func(http.ResponseWriter, *http.Request)); ok {
			route := router.HandleFunc(field.Path, handlerFunc)
			if field.Method == mir.MethodAny {
				route.Methods([]string{
					mir.MethodGet,
					mir.MethodPut,
					mir.MethodPost,
					mir.MethodDelete,
					mir.MethodHead,
					mir.MethodPatch,
					mir.MethodOptions,
					mir.MethodConnect,
					mir.MethodTrace,
				}...)
			} else {
				route.Methods(field.Method)
			}
			if err := inflateQueries(route, field.Queries); err != nil {
				return err
			}
			if field.Host != "" {
				route.Host(field.Host)
			}
		} else {
			return fmt.Errorf("handler not func(http.ResponseWriter, *http.Request) function")
		}
	}
	return nil
}

// handlerChainTo setup handlers to router that grouped
func handlerChainTo(router *mux.Router, chain mir.Chain) error {
	// just return if empty chain
	if chain == nil {
		return nil
	}
	if handlerChain, ok := chain.([]mux.MiddlewareFunc); ok {
		router.Use(handlerChain...)
	} else {
		return fmt.Errorf("chain type not []mux.MiddlewareFunc")
	}
	return nil
}

// inflateQueries setup queries to route
func inflateQueries(route *mux.Route, queries []string) error {
	if len(queries) == 0 {
		return nil
	}

	// length will x2 of queries
	fixedQueries := make([]string, 0, len(queries)*2)

	// start inflate queries
	for _, query := range queries {
		kv := strings.Split(query, "=")
		if len(kv) != 2 {
			return fmt.Errorf("mir.mux: query must like ?filter={filter}&foo=bar")
		}
		fixedQueries = append(fixedQueries, kv...)
	}

	// setup fixed queries to route
	route.Queries(fixedQueries...)
	return nil
}
