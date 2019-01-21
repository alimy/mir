// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package iris

import (
	"fmt"
	"github.com/alimy/mir"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

var _ mir.Engine = &mirEngine{}

// mirEngine used to implements mir.Engine interface
type mirEngine struct {
	engine *iris.Application
}

// Register register entries to echo engine
func (e *mirEngine) Register(entries []*mir.TagMir) error {
	for _, entry := range entries {
		var party iris.Party
		if entry.Group == "" || entry.Group == "/" {
			party = e.engine
		} else {
			party = e.engine.Party(entry.Group)
		}
		if err := handlerChainTo(party, entry.Chain); err != nil {
			return err
		}
		// Notice just return if catch a error or continue next entry register
		if err := registerWith(party, entry.Fields); err != nil {
			return err
		}
	}
	return nil
}

// registerWith register fields to give router
func registerWith(party iris.Party, fields []*mir.TagField) error {
	for _, field := range fields {
		if handlerFunc, ok := field.Handler.(func(context.Context)); ok {
			if field.Method == mir.MethodAny {
				party.Any(field.Path, handlerFunc)
			} else {
				party.Handle(field.Method, field.Path, handlerFunc)
			}
		} else {
			return fmt.Errorf("handler not function of func(context.Context)")
		}
	}
	return nil
}

// handlerChainTo setup handlers to router that grouped
func handlerChainTo(party iris.Party, chain mir.Chain) error {
	// just return if empty chain
	if chain == nil {
		return nil
	}
	if handlerChain, ok := chain.(context.Handlers); ok {
		party.Use(handlerChain...)
	} else {
		return fmt.Errorf("chain type not context.Handlers")
	}
	return nil
}
