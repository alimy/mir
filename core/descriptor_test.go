// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package core

import (
	"testing"

	"github.com/alimy/mir/v3"
)

func TestDescriptors(t *testing.T) {
	d := make(Descriptors)
	if err := d.Put(&IfaceDescriptor{
		Group:    "",
		Chain:    "Chain",
		PkgName:  "api",
		TypeName: "site",
		Comment:  "",
		Fields: []*FieldDescriptor{
			{
				Host:       "",
				Path:       "/",
				Queries:    nil,
				HttpMethod: mir.MethodGet,
				MethodName: "Index",
				Comment:    "",
			},
		},
	}); err != nil {
		t.Error("want nil error but not")
	}
	if _, exist := d.Get(""); !exist {
		t.Error("want exist an item but not")
	}
}
