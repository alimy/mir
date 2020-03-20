package internal

import (
	"errors"
	"testing"
)

func TestMirCtx(t *testing.T) {
	ctx := NewMirCtx(10)
	if ctx.Capcity() != 10 {
		t.Error("want capcity is 10 but not")
	}
	ctx.Cancel(errors.New("just an error"))
	ctx.Cancel(errors.New("another error"))
	if err := ctx.Err(); err != nil && err.Error() != "just an error" {
		t.Error("want 'just an error' but not")
	}
}
