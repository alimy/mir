// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package utils

import "testing"

func TestString(t *testing.T) {
	for _, d := range []struct {
		input  []byte
		expect string
	}{
		{[]byte(""), ""},
		{[]byte("sail-001"), "sail-001"},
		{[]byte("sail-002"), "sail-002"},
		{[]byte("sail-003"), "sail-003"},
	} {
		if out := String(d.input); out != d.expect {
			t.Errorf("give %s want %s but got %s", string(d.input), d.expect, out)
		}
	}
}

func TestBytes(t *testing.T) {
	for _, d := range []struct {
		input  string
		expect []byte
	}{
		{"", []byte{}},
		{"sail-001", []byte("sail-001")},
		{"sail-002", []byte("sail-002")},
		{"sail-003", []byte("sail-003")},
	} {
		if out := Bytes(d.input); string(out) != string(d.expect) {
			t.Errorf("give %s want %s but got %s", d.input, string(d.expect), out)
		}
	}
}
