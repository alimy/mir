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
		// notice just return if catch a error
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
