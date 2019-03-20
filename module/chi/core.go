// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package chi

import (
	"fmt"
	"github.com/alimy/mir"
	"github.com/go-chi/chi"
	"net/http"
	"strings"
)

var _ mir.Engine = &mirEngine{}

// mirEngine used to implements mir.Engine interface
type mirEngine struct {
	engine chi.Router
}

// Register register entries to chi engine
func (e *mirEngine) Register(entries []*mir.TagMir) error {
	for _, entry := range entries {
		var router chi.Router
		if entry.Group == "" || entry.Group == "/" {
			router = e.engine
		} else {
			pathPrefix := entry.Group
			if !strings.HasPrefix(entry.Group, "/") {
				pathPrefix = "/" + entry.Group
			}
			router = chi.NewRouter()
			e.engine.Mount(pathPrefix, router)
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
func registerWith(router chi.Router, fields []*mir.TagField) error {
	for _, field := range fields {
		if handlerFunc, ok := field.Handler.(func(http.ResponseWriter, *http.Request)); ok {
			if field.Method == mir.MethodAny {
				router.Connect(field.Path, handlerFunc)
				router.Delete(field.Path, handlerFunc)
				router.Get(field.Path, handlerFunc)
				router.Head(field.Path, handlerFunc)
				router.Options(field.Path, handlerFunc)
				router.Patch(field.Path, handlerFunc)
				router.Post(field.Path, handlerFunc)
				router.Put(field.Path, handlerFunc)
				router.Trace(field.Path, handlerFunc)
			} else {
				router.MethodFunc(field.Method, field.Path, handlerFunc)
			}
		} else {
			return fmt.Errorf("handler not func(http.ResponseWriter, *http.Request) function")
		}
	}
	return nil
}

// handlerChainTo setup handlers to router that grouped
func handlerChainTo(router chi.Router, chain mir.Chain) error {
	// just return if empty chain
	if chain == nil {
		return nil
	}
	if middlewares, ok := chain.(chi.Middlewares); ok {
		router.Use(middlewares...)
	} else {
		return fmt.Errorf("chain type not chi.Middlewares")
	}
	return nil
}
