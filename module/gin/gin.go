package gin

import (
	"github.com/alimy/mir"
	"github.com/gin-gonic/gin"
)

func Mir(e *gin.Engine) mir.Engine {
	return &MirEngine{Engine: e}
}
