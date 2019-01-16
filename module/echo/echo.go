package echo

import (
	"github.com/alimy/mir"
	"github.com/labstack/echo"
)

// Mir return mir.Engine interface implements instance.Used to register routes
// to gin engine with struct tag string's information.
func Mir(e *echo.Echo) mir.Engine {
	return &mirEngine{engine: e}
}
