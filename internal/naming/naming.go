// Copyright 2025 Michael Li <alimy@niubiu.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package naming

// NamingStrategy naming strategy interface
type NamingStrategy interface {
	Naming(string) string
}
