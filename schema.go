// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package mir

type (
	Interface interface {
		mustEmbedSchema()
	}
	// Schema is the default implementation for the restful Interface.
	// It can be embedded in end-user schemas as follows:
	//
	//	type T struct {
	//		mir.Schema
	//	}
	//
	Schema struct{}
)

func (Schema) mustEmbedSchema() {}
