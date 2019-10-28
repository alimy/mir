// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package httprouter

import (
	"fmt"
	"net/http"

	"github.com/alimy/mir"
	"github.com/julienschmidt/httprouter"
)

var _ mir.Engine = &mirEngine{}

// mirEngine used to implements mir.Engine interface
type mirEngine struct {
	engine *httprouter.Router
}

// Register register entries to echo engine
func (e *mirEngine) Register(entries []*mir.TagMir) error {
	for _, entry := range entries {
		for _, field := range entry.Fields {
			if handlerFunc, ok := field.Handler.(func(http.ResponseWriter, *http.Request, httprouter.Params)); ok {
				if field.Method == mir.MethodAny {
					e.engine.Handle(mir.MethodGet, field.Path, handlerFunc)
					e.engine.Handle(mir.MethodPut, field.Path, handlerFunc)
					e.engine.Handle(mir.MethodPost, field.Path, handlerFunc)
					e.engine.Handle(mir.MethodDelete, field.Path, handlerFunc)
					e.engine.Handle(mir.MethodHead, field.Path, handlerFunc)
					e.engine.Handle(mir.MethodPatch, field.Path, handlerFunc)
					e.engine.Handle(mir.MethodOptions, field.Path, handlerFunc)
					e.engine.Handle(mir.MethodConnect, field.Path, handlerFunc)
					e.engine.Handle(mir.MethodTrace, field.Path, handlerFunc)

				} else {
					e.engine.Handle(field.Method, field.Path, handlerFunc)
				}
			} else {
				return fmt.Errorf("handler not function of func(http.ResponseWriter, *http.Request, httprouter.Params)")
			}
		}
	}
	return nil
}
