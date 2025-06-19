// Copyright 2025 Michael Li <alimy@niubiu.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package engine

import (
	"testing"

	"github.com/alimy/mir/v5/core"
)

func TestGenerate(t *testing.T) {
	if err := Generate(core.WithEntry(nil)); err == nil {
		t.Error("want an error but not")
	}
	AddEntry(nil)
	AddEntry(nil, nil)
	if len(mirEntries) != 4 {
		t.Errorf("want mirEntries's size is 4 but is %d", len(mirEntries))
	}
}
