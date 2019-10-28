// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package macaron

import (
	"fmt"

	"github.com/alimy/mir"
	"gopkg.in/macaron.v1"
)

var _ mir.Engine = &mirEngine{}

// mirEngine used to implements mir.Engine interface
type mirEngine struct {
	engine *macaron.Macaron
}

// Register register entries to Macaron engine
func (e *mirEngine) Register(entries []*mir.TagMir) error {
	for _, entry := range entries {
		if entry.Group == "" || entry.Group == "/" {
			handlers, err := e.handlerChain(entry.Chain)
			if err != nil {
				return err
			}
			for _, handler := range handlers {
				e.engine.Use(handler)
			}
			e.registerWith(entry.Fields)
		} else {
			if err := e.registerGroup(entry); err != nil {
				return err
			}
		}
	}
	return nil
}

// registerGroup register entry by group
func (e *mirEngine) registerGroup(entry *mir.TagMir) error {
	handlers, err := e.handlerChain(entry.Chain)
	if err != nil {
		return err
	}
	e.engine.Group(entry.Group, func() {
		e.registerWith(entry.Fields)
	}, handlers...)
	return nil
}

// registerWith register fields engine
func (e *mirEngine) registerWith(fields []*mir.TagField) {
	for _, field := range fields {
		if field.Method == mir.MethodAny {
			e.engine.Any(field.Path, []macaron.Handler{field.Handler}...)
		} else {
			e.engine.Handle(field.Method, field.Path, []macaron.Handler{field.Handler})
		}
	}
}

// handlerChain get handlers to router that grouped
func (e *mirEngine) handlerChain(chain mir.Chain) ([]macaron.Handler, error) {
	// just return if empty chain
	if chain == nil {
		return []macaron.Handler{}, nil
	}
	if handlerChain, ok := chain.([]macaron.Handler); ok {
		return handlerChain, nil
	}
	return []macaron.Handler{}, fmt.Errorf("chain type not macaron.Handler")
}
