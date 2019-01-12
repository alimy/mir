package gin

import (
	"fmt"
	"github.com/alimy/mir"
	"github.com/gin-gonic/gin"
)

// mirEngine used to implements mir.Engine interface
type mirEngine struct {
	engine *gin.Engine
}

// Register register entries to gin engine
func (e *mirEngine) Register(entries ...interface{}) error {
	handlerGroup := make(map[string][]*mir.TagField)

	// collect TagField by group info
	for _, entry := range entries {
		if tagFields, err := mir.TagFieldsFrom(entry); err == nil {
			for _, tagField := range tagFields {
				fileds := handlerGroup[tagField.Group]
				if fileds == nil {
					fileds = make([]*mir.TagField, 0)
				}
				handlerGroup[tagField.Group] = append(fileds, tagField)
			}
		} else {
			return err
		}
	}

	// register TagFields from handlerGroup
	for group, fields := range handlerGroup {
		var router gin.IRouter
		if group == "" || group == "/" {
			router = e.engine
		} else {
			router = e.engine.Group(group)
		}
		if err := registerWith(router, fields); err != nil {
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
}
