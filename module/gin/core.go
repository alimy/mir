package gin

import "github.com/gin-gonic/gin"

type MirEngine struct {
	Engine *gin.Engine
}

func (e *MirEngine) Register(entries ...interface{}) error {
	// TODO
}
