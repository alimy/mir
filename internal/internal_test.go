// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

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
