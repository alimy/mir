// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package faygo

import (
	"fmt"
	"github.com/alimy/mir"
	"github.com/henrylee2cn/faygo"
)

var _ mir.Engine = &mirEngine{}

// mirEngine used to implements mir.Engine interface
type mirEngine struct {
	engine *faygo.Framework
}

// Register register entries to gin engine
func (e *mirEngine) Register(entries []*mir.TagMir) error {
	for _, entry := range entries {
		var mux *faygo.MuxAPI
		if entry.Group == "" || entry.Group == "/" {
			mux = e.engine.MuxAPI
		} else {
			mux = e.engine.Group(entry.Group)
		}
		if err := handlerChainTo(mux, entry.Chain); err != nil {
			return err
		}
		// Notice just return if catch a error or continue next entry register
		if err := registerWith(mux, entry.Fields); err != nil {
			return err
		}
	}
	return nil
}

// registerWith register fields to give mux
func registerWith(mux *faygo.MuxAPI, fields []*mir.TagField) error {
	for _, field := range fields {
		if handler, ok := field.Handler.(faygo.Handler); ok {
			if field.Method == mir.MethodAny {
				for _, method := range faygo.RESTfulMethodList {
					mux.API(faygo.Methodset(method), field.Path, handler)
				}
			} else {
				mux.API(faygo.Methodset(field.Method), field.Path, handler)
			}
		} else {
			return fmt.Errorf("handler not faygo.Handler type")
		}
	}
	return nil
}

// handlerChainTo setup handlers to mux that grouped
func handlerChainTo(mux *faygo.MuxAPI, chain mir.Chain) error {
	// just return if empty chain
	if chain == nil {
		return nil
	}
	if handlerChain, ok := chain.([]faygo.Handler); ok {
		mux.Use(handlerChain...)
	} else {
		return fmt.Errorf("chain type not []faygo.Handler type")
	}
	return nil
}
