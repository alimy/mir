package engine

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	if err := Generate(nil, nil); err != nil {
		t.Error("don't want an error but not")
	}
	AddEntry(nil)
	AddEntries(nil, nil)
	if len(mirEntries) != 4 {
		t.Errorf("want mirEntries's size is 4 but is %d", len(mirEntries))
	}
}
