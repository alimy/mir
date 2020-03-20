package engine

import (
	"testing"

	"github.com/alimy/mir/v2/core"
)

func TestGenerate(t *testing.T) {
	if err := Generate(nil, nil); err == nil {
		t.Error("want an error but not")
	}
	if err := Generate(nil, core.DefaultOptions()); err == nil {
		t.Error("want an error but not")
	}
}
