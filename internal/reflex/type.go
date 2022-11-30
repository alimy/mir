// Copyright 2022 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package reflex

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"

	"github.com/alimy/mir/v3/internal/utils"
)

// CheckStruct check struct type is in pkgPath and return other struct type field's
// type that contained in type.
// st must struct kind
func CheckStruct(st reflect.Type, pkgPath string) ([]reflect.Type, error) {
	sts := []reflect.Type{st}
	fields := utils.NewStrSet()
	fields.Add(st.PkgPath() + "." + st.Name())
	for i := 0; i < len(sts); i++ {
		st := sts[i]
		if st.PkgPath() != pkgPath {
			return nil, errors.New("pkgPath need in same path")
		}
		for i := st.NumField() - 1; i >= 0; i-- {
			sf := st.Field(i)
			ft := sf.Type
			ok, wbsts := isWriteableField(ft)
			if !ok {
				return nil, fmt.Errorf("yet not support field %v", ft)
			}
			for _, t := range wbsts {
				if t.PkgPath() != pkgPath {
					return nil, errors.New("struct field must in same path")
				}
				fn := t.PkgPath() + "." + ft.Name()
				if err := fields.Add(fn); err == nil {
					sts = append(sts, t)
				}
			}
		}
	}
	return sts, nil
}

// WriteStruct write struct type to *bytes.Buffer
func WriteStruct(buf *bytes.Buffer, t reflect.Type, indent string) (err error) {
	if buf == nil || t == nil {
		return errors.New("buf or t of type is nil")
	}
	if _, err = buf.WriteString(fmt.Sprintf("type %s struct {\n", t.Name())); err != nil {
		return
	}
	l := t.NumField()
	for i := 0; i < l; i++ {
		field := t.Field(i)
		if err = writeStructField(buf, field, indent); err != nil {
			return
		}
	}
	_, err = buf.WriteString("}\n")
	return
}

func writeStructField(buf *bytes.Buffer, sf reflect.StructField, indent string) (err error) {
	if ok, _ := isWriteableField(sf.Type); !ok {
		return
	}
	buf.WriteString(indent)
	if !sf.Anonymous {
		if _, err = buf.WriteString(fmt.Sprintf("%s%s", sf.Name, indent)); err != nil {
			return
		}
	}
	if err = writeStructFieldType(buf, sf.Type); err != nil {
		return
	}
	if len(sf.Tag) > 0 {
		_, err = buf.WriteString(fmt.Sprintf("%s`%s`\n", indent, sf.Tag))
	} else {
		_, err = buf.WriteString("\n")
	}
	return
}

func writeStructFieldType(buf *bytes.Buffer, t reflect.Type) (err error) {
	ft := t
	for ; ft.Kind() == reflect.Pointer; ft = ft.Elem() {
		if err = buf.WriteByte('*'); err != nil {
			return
		}
	}
	switch ft.Kind() {
	case reflect.String, reflect.Bool,
		reflect.Complex64, reflect.Complex128,
		reflect.Float32, reflect.Float64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if _, err = buf.WriteString(ft.String()); err != nil {
			return
		}
	case reflect.Array:
		_, err = buf.WriteString(fmt.Sprintf("[%d]", ft.Len()))
		if err != nil {
			return
		}
		err = writeStructFieldType(buf, ft.Elem())
		if err != nil {
			return
		}
	case reflect.Slice:
		if _, err = buf.WriteString("[]"); err != nil {
			return
		}
		if err = writeStructFieldType(buf, ft.Elem()); err != nil {
			return
		}
	case reflect.Map:
		if _, err = buf.WriteString("map["); err != nil {
			return
		}
		if err = writeStructFieldType(buf, ft.Key()); err != nil {
			return
		}
		if _, err = buf.WriteString("]"); err != nil {
			return
		}
		if err = writeStructFieldType(buf, ft.Elem()); err != nil {
			return
		}
	case reflect.Struct:
		if _, err = buf.WriteString(ft.Name()); err != nil {
			return
		}
	}
	return
}

func isWriteableField(t reflect.Type) (b bool, sts []reflect.Type) {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	switch t.Kind() {
	case reflect.String, reflect.Bool,
		reflect.Complex64, reflect.Complex128,
		reflect.Float32, reflect.Float64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true, nil
	case reflect.Array, reflect.Slice:
		t = t.Elem()
		return isWriteableField(t)
	case reflect.Map:
		kb, kt := isWriteableMapKey(t.Key())
		vb, vt := isWriteableField(t.Elem())
		if kt != nil {
			sts = append(sts, kt)
		}
		sts = append(sts, vt...)
		b = kb && vb
		return
	case reflect.Struct:
		b = len(t.Name()) > 0
		sts = append(sts, t)
		return
	}
	return false, nil
}

func isWriteableMapKey(t reflect.Type) (bool, reflect.Type) {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	switch t.Kind() {
	case reflect.String, reflect.Bool,
		reflect.Complex64, reflect.Complex128,
		reflect.Float32, reflect.Float64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true, nil
	case reflect.Struct:
		return len(t.Name()) > 0, t
	}
	return false, nil
}
