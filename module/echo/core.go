// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package echo

import (
	"fmt"
	"github.com/alimy/mir"
	"github.com/labstack/echo"
)

var _ mir.Engine = &mirEngine{}

// mirEngine used to implements mir.Engine interface
type mirEngine struct {
	engine *echo.Echo
}

// Register register entries to echo engine
func (e *mirEngine) Register(entries []*mir.TagMir) error {
	for _, entry := range entries {
		g := e.engine.Group(entry.Group)
		if err := handlerChainTo(g, entry.Chain); err != nil {
			return err
		}
		// Notice just return if catch a error or continue next entry register
		if err := registerWith(g, entry.Fields); err != nil {
			return err
		}
	}
	return nil
}

// registerWith register fields to give router
func registerWith(g *echo.Group, fields []*mir.TagField) error {
	for _, field := range fields {
		if handlerFunc, ok := field.Handler.(func(echo.Context) error); ok {
			if field.Method == mir.MethodAny {
				g.Any(field.Path, handlerFunc)
			} else {
				g.Add(field.Method, field.Method, handlerFunc)
			}
		} else {
			return fmt.Errorf("handler not function of func(echo.Context) error")
		}
	}
	return nil
}

// handlerChainTo setup handlers to router that grouped
func handlerChainTo(g *echo.Group, chain mir.Chain) error {
	// just return if empty chain
	if chain == nil {
		return nil
	}
	if handlerChain, ok := chain.([]echo.MiddlewareFunc); ok {
		g.Use(handlerChain...)
	} else {
		return fmt.Errorf("chain type not []echo.MiddlewareFunc")
	}
	return nil
}
