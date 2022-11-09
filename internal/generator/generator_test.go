// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package generator

import (
	"testing"

	"github.com/alimy/mir/v3/core"
)

func TestMirGenerator_Init(t *testing.T) {
	g := &mirGenerator{name: core.GeneratorGin}
	if err := g.Init(nil); err == nil {
		t.Error("want an error but not")
	}
}
