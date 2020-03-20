package generator

import (
	"testing"

	"github.com/alimy/mir/v2/core"
)

func TestMirGenerator_Init(t *testing.T) {
	g := &mirGenerator{name: core.GeneratorGin}
	if err := g.Init(nil); err == nil {
		t.Error("want an error but not")
	}
}
