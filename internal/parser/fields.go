// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package parser

import (
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/alimy/mir/v2"
)

var (
	// error list
	errNotExist      tagError = "mir struct tag filed not exist"
	errNoPathInfo    tagError = "mir struct tag not contains path info"
	errNotValideType tagError = "not valide type, just struct and struct ptr is avalibale"
	errMultGroupInfo tagError = "multiple group info in struct defined"
	errMultChainInfo tagError = "multiple chain info in struct defined"

	// defaultTag indicate default mir's struct tag name
	defaultTag = "mir"
)

// tagError indicate error information
type tagError string

// Error error message string
func (e tagError) Error() string {
	return string(e)
}

// tagInfo indicate mir tag information in struct tag string
type tagInfo struct {
	Method    string   // Method indicate method information in struct tag string
	Host      string   // Host indicate host information in struct tag string
	Path      string   // Path indicate path information in struct tag string
	Queries   []string // Queries indicate path information in struct tag string
	isGroup   bool     // indicate whether a group field
	isChain   bool     // indicate whether a chain field
	group     string   // indicate group information in struct tag string
	chainFunc string   // indicate chain function information in struct tag string
	handler   string   // indicate handler information in struct tag string
	fieldName string   // indicate field name
	comment   string   // indicate comment info (not support now)
}

// tagInfoFrom build tagInfo from field
func tagInfoFrom(field reflect.StructField) (*tagInfo, error) {
	info := &tagInfo{}

	// lookup mir tag info from struct field
	tag, exist := field.Tag.Lookup(defaultTag)
	if !exist {
		return nil, errNotExist
	}

	// Skip leading space.
	i := 0
	for i < len(tag) && tag[i] == ' ' {
		i++
	}
	tag = tag[i:]

	// group info or method info or chain info
	info.fieldName = field.Name
	switch field.Type.Name() {
	case "Chain":
		info.isChain = true
		return info, nil
	case "Group":
		info.isGroup = true
		info.group = tag
		return info, nil
	case "Get":
		info.Method = mir.MethodGet
	case "Put":
		info.Method = mir.MethodPut
	case "Post":
		info.Method = mir.MethodPost
	case "Delete":
		info.Method = mir.MethodDelete
	case "Head":
		info.Method = mir.MethodHead
	case "Options":
		info.Method = mir.MethodOptions
	case "Patch":
		info.Method = mir.MethodPatch
	case "Trace":
		info.Method = mir.MethodTrace
	case "Connect":
		info.Method = mir.MethodConnect
	case "Any":
		info.Method = "ANY"
	}

	// host info
	if len(tag) > 2 && tag[0] == '/' && tag[1] == '/' {
		i := 2
		for i < len(tag) && tag[i] != '/' {
			i++
		}
		info.Host = tag[2:i]
		tag = tag[i:]
	}

	// path info. must have path info if not a group field
	if len(tag) == 0 && !info.isGroup {
		return nil, errNoPathInfo
	}
	i = 0
	for i < len(tag) && tag[i] != '?' && tag[i] != '#' {
		i++
	}
	info.Path = tag[0:i]
	tag = tag[i:]

	// queries and handler info
	for len(tag) != 0 {
		switch tag[0] {
		case '#':
			i := 1
			for i < len(tag) && tag[i] != '?' {
				i++
			}
			handlerStr := tag[1:i]
			tag = tag[i:]
			if handlerStr != "" {
				if handlerStr[0] == '-' { // just contain chain func info
					info.chainFunc = handlerStr[1:]
				} else { // contain handler and inline chain info like #Handler&ChainFunc
					handlerChains := strings.Split(handlerStr, "&")
					info.handler = handlerChains[0]
					if len(handlerChains) > 1 { // extract chain func
						info.chainFunc = handlerChains[1]
					}
				}
			}
		case '?':
			i := 1
			for i < len(tag) && tag[i] != '#' {
				i++
			}
			queryStr := tag[1:i]
			if queryStr != "" {
				info.Queries = inflateQuery(queryStr)
			}
			tag = tag[i:]
		}
	}

	// check handler if not group field
	if info.handler == "" {
		firstRune, size := utf8.DecodeRuneInString(field.Name)
		upperFirst := unicode.ToUpper(firstRune)

		// encode upperFirst to []byte,use max byte for contain unicode
		methoName := make([]byte, 4)
		number := utf8.EncodeRune(methoName, upperFirst)
		methoName = methoName[:number]
		methoName = append(methoName, field.Name[size:]...)

		// assign handler name
		info.handler = string(methoName)
	}

	return info, nil
}

func inflateQuery(qs string) []string {
	items := strings.Split(qs, "&")
	res := make([]string, 0, len(items)*2)
	for _, q := range items {
		kv := strings.Split(q, "=")
		if len(kv) == 2 {
			res = append(res, kv...)
		}
	}
	return res
}

// valueByName return field value by field name
func valueByName(value reflect.Value, name string) interface{} {
	if fieldValue := value.FieldByName(name); !fieldValue.IsNil() {
		return fieldValue.Elem().Interface()
	}
	return nil
}
