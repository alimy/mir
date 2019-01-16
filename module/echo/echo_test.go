package echo

import (
	"github.com/alimy/mir"
	"github.com/labstack/echo"
	"net/http"
	"testing"
)

type site struct {
	count    uint64
	v1       mir.Group `mir:"v1"`
	add      mir.Post  `mir:"/add/:id"`
	index    mir.Get   `mir:"/index/"`
	articles mir.Get   `mir:"/articles/:category/#GetArticles"`
}

// Add handler of "/add/:id"
func (h *site) Add(c echo.Context) error {
	return c.String(http.StatusOK, "add")
}

// Index handler of the index field that in site struct, the struct tag indicate
// this handler will register to path "/index/" and method is http.MethodGet.
func (h *site) Index(c echo.Context) error {
	h.count++
	return c.String(http.StatusOK, "Index")
}

// GetArticles handler of articles indicator that contains Host/Path/Queries/Handler info.
func (h *site) GetArticles(c echo.Context) error {
	return c.String(http.StatusOK, "GetArticles")
}

func TestMir(t *testing.T) {
	e := echo.New()
	mirE := Mir(e)
	if err := mir.Register(mirE, &site{}); err != nil {
		t.Error(err)
	}
	// TODO: add httptest assert
}
