package engine

import (
	"testing"

	"github.com/alimy/mir/v2/core"
)

func TestGenerate(t *testing.T) {
	if err := Generate(nil); err == nil {
		t.Error("want an error but not")
	}
	if err := Generate(core.DefaultOptions()); err == nil {
		t.Error("want an error but not")
	}
	if err := Generate(core.DefaultOptions(), nil); err != nil {
		t.Error("don't want an error but not")
	}
	AddEntry(nil)
	AddEntries(nil, nil)
	if len(mirEntries) != 4 {
		t.Errorf("want mirEntries's size is 4 but is %d", len(mirEntries))
	}
}
