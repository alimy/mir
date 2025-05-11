// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package utils

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// UpperFirst make first rune upper in s string
func UpperFirst(s string) string {
	firstRune, size := utf8.DecodeRuneInString(s)
	if unicode.IsUpper(firstRune) {
		return s
	}
	// encode upperFirst to []byte,use max byte for contain unicode
	res := make([]byte, len(s))
	upperRune := unicode.ToUpper(firstRune)
	number := utf8.EncodeRune(res, upperRune)
	res = res[:number]
	res = append(res, s[size:]...)
	return string(res)
}

// QuoteJoin give a string slice return a quoted join string by sep
func QuoteJoin(elems []string, sep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(elems[i]) + 2
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteByte('"')
	b.WriteString(elems[0])
	b.WriteByte('"')
	for _, s := range elems[1:] {
		b.WriteString(sep)
		b.WriteByte('"')
		b.WriteString(s)
		b.WriteByte('"')
	}
	return b.String()
}
