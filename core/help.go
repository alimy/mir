package core

func AssertTypeSpec(style string) (spec string, imports []string) {
	switch style {
	case GeneratorChi:
		spec, imports = "AssertType[*chi.Context]", []string{"github.com/go-chi/chi/v5"}
	case GeneratorMux:
		spec, imports = "AssertType2[*http.Request, http.ResponseWriter]", []string{"net/http"}
	case GeneratorHertz:
		spec, imports = "AssertType4[context.Context, *app.RequestContext]", []string{"context", "github.com/cloudwego/hertz/pkg/app"}
	case GeneratorEcho:
		spec, imports = "AssertType[echo.Context]", []string{"github.com/labstack/echo/v4"}
	case GeneratorIris:
		spec, imports = "AssertType[*context.Context]", []string{"github.com/kataras/iris/v12/context"}
	case GeneratorFiber:
		spec, imports = "AssertType[*fiber.Ctx]", []string{"github.com/gofiber/fiber/v2"}
	case GeneratorMacaron:
		spec, imports = "AssertType[*macaron.Context]", []string{"gopkg.in/macaron.v1"}
	case GeneratorHttpRouter:
		spec, imports = "AssertType3[*http.Request, httprouter.Params, http.ResponseWriter]", []string{"net/http", "github.com/julienschmidt/httprouter"}
	case GeneratorGin:
		fallthrough
	default:
		spec, imports = "AssertType[*gin.Context]", []string{"github.com/gin-gonic/gin"}
	}
	return
}
