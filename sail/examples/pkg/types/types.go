// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package types

const (
	// No  二态值 否
	No int8 = 0

	// Yes  二态值 是
	Yes int8 = 1
)

// Empty empty alias type
type Empty = struct{}

// Fn empty argument func alias type
type Fn = func()

// Boxes Box/Unbox interface
type Boxes[T any] interface {
	Box(t T)
	Unbox() T
}

type Integer interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int | ~uint
}
