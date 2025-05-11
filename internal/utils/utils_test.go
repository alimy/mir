// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package utils

import "testing"

func TestUpperFirst(t *testing.T) {
	for input, expect := range map[string]string{
		"Options": "Options",
		"get":     "Get",
		"post":    "Post",
		"head":    "Head",
	} {
		if res := UpperFirst(input); res != expect {
			t.Errorf("expect %s but got %s", expect, res)
		}
	}
}

func TestQuoteJoin(t *testing.T) {
	for _, data := range []struct {
		input  []string
		expect string
	}{
		{
			input:  []string{"GET", "POST", "HEAD"},
			expect: `"GET","POST","HEAD"`,
		},
		{
			input:  []string{"Options", "Trace", "Post"},
			expect: `"Options","Trace","Post"`,
		},
	} {
		if res := QuoteJoin(data.input, ","); res != data.expect {
			t.Errorf("expect (%s) but got (%s)", data.expect, res)
		}
	}
}
