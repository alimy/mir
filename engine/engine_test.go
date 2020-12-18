// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

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
