package gin

import (
	"github.com/alimy/mir"
	"github.com/gin-gonic/gin"
)

// Mir return mir.Engine interface implements instance.Used to register routes
// to gin engine with struct tag string's information.
func Mir(e *gin.Engine) mir.Engine {
	return &mirEngine{engine: e}
}
