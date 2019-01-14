// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package gin

import (
	"fmt"
	"github.com/alimy/mir"
	"github.com/gin-gonic/gin"
)

var _ mir.Engine = &mirEngine{}

// mirEngine used to implements mir.Engine interface
type mirEngine struct {
	engine *gin.Engine
}

// Register register entries to gin engine
func (e *mirEngine) Register(entries []*mir.TagMir) error {
	for _, entry := range entries {
		var router gin.IRouter
		if entry.Group == "" || entry.Group == "/" {
			router = e.engine
		} else {
			router = e.engine.Group(entry.Group)
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
func registerWith(router gin.IRouter, fields []*mir.TagField) error {
	for _, field := range fields {
		if handlerFunc, ok := field.Handler.(func(*gin.Context)); ok {
			if field.Method == mir.MethodAny {
				router.Any(field.Path, handlerFunc)
			} else {
				router.Handle(field.Method, field.Path, handlerFunc)
			}
		} else {
			return fmt.Errorf("handler not func(*gin.Context) function")
		}
	}
	return nil
}

// handlerChainTo setup handlers to router that grouped
func handlerChainTo(router gin.IRouter, chain mir.Chain) error {
	// just return if empty chain
	if chain == nil {
		return nil
	}
	if handlerChain, ok := chain.(gin.HandlersChain); ok {
		router.Use(handlerChain...)
	} else {
		return fmt.Errorf("chain type not gin.HandlersChain")
	}
	return nil
}
