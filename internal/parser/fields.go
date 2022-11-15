// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package parser

import (
	"errors"
	"github.com/alimy/mir/v3"
	"github.com/alimy/mir/v3/internal/container"
	"net/http"
	"reflect"
	"strings"

	"github.com/alimy/mir/v3/internal/utils"
)

var (
	// error list
	errNilType       tagError = "nil type is not valide"
	errNotExist      tagError = "mir struct tag filed not exist"
	errNoPathInfo    tagError = "mir struct tag not contains path info"
	errNotValideType tagError = "not valide type, just struct and struct ptr is avalibale"
	errMultGroupInfo tagError = "multiple group info in struct defined"
	errMultChainInfo tagError = "multiple chain info in struct defined"

	// defaultTag indicate default mir's struct tag name
	defaultTag       = "mir"
	defautlMethodTag = "method"
)

// tagError indicate error information
type tagError string

// Error error message string
func (e tagError) Error() string {
	return string(e)
}

// tagInfo indicate mir tag information in struct tag string
type tagInfo struct {
	Methods   container.HttpMethodSet // Method indicate methods information in struct tag string
	Host      string                  // Host indicate host information in struct tag string
	Path      string                  // Path indicate path information in struct tag string
	Queries   []string                // Queries indicate path information in struct tag string
	isGroup   bool                    // indicate whether a group field
	isChain   bool                    // indicate whether a chain field
	group     string                  // indicate group information in struct tag string
	chainFunc string                  // indicate chain function information in struct tag string
	handler   string                  // indicate handler information in struct tag string
	fieldName string                  // indicate field name
	in        reflect.Type
	out       reflect.Type
	inOuts    []reflect.Type
	comment   string // indicate comment info (not support now)
}

// tagInfoFrom build tagInfo from field
func (r *reflex) tagInfoFrom(field reflect.StructField, pkgPath string) (*tagInfo, error) {
	info := &tagInfo{
		Methods: make(container.HttpMethodSet, 1),
	}

	// lookup mir tag info from struct field
	tag, exist := field.Tag.Lookup(r.tagName)
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
	switch field.Type.Kind() {
	case reflect.Interface:
		if field.Type.PkgPath() != "github.com/alimy/mir/v3" {
			return nil, errors.New("not supported filed type")
		}
		switch field.Type.Name() {
		case "Chain":
			info.isChain = true
			return info, nil
		case "Group":
			info.isGroup = true
			info.group = tag
			return info, nil
		default:
			return nil, errors.New("not supported filed type")
		}
	case reflect.Func:
		ft := field.Type
		numIn := ft.NumIn()
		numOut := ft.NumOut()
		if numIn == 0 || numOut > 1 {
			return nil, errors.New("func field must need one least argument and one most return value")
		}

		// request type in latest in argument if declared
		it := ft.In(numIn - 1)
		if it.Kind() == reflect.Struct {
			cts, err := checkStruct(it, pkgPath)
			if err != nil {
				return nil, err
			}
			info.in = it
			info.inOuts = append(info.inOuts, cts...)

			// minus numIn to ignore latest in argument that had processed
			numIn--
		}

		// process other in argument
		for i := numIn - 1; i >= 0; i-- {
			it = ft.In(i)
			if it.PkgPath() != "github.com/alimy/mir/v3" {
				continue
			}
			switch it.Name() {
			case "Get":
				info.Methods.Add(mir.MethodGet)
			case "Put":
				info.Methods.Add(mir.MethodPut)
			case "Post":
				info.Methods.Add(mir.MethodPost)
			case "Delete":
				info.Methods.Add(mir.MethodDelete)
			case "Head":
				info.Methods.Add(mir.MethodHead)
			case "Options":
				info.Methods.Add(mir.MethodOptions)
			case "Patch":
				info.Methods.Add(mir.MethodPatch)
			case "Trace":
				info.Methods.Add(mir.MethodTrace)
			case "Connect":
				info.Methods.Add(mir.MethodConnect)
			case "Any":
				info.Methods.Add(mir.HttpMethods...)
			}
		}
		if numOut == 1 {
			ot := ft.Out(i)
			if ot.Kind() != reflect.Struct {
				return nil, errors.New("func field must return value is need struct type")
			}
			cts, err := checkStruct(ot, pkgPath)
			if err != nil {
				return nil, err
			}
			info.out = ot
			info.inOuts = append(info.inOuts, cts...)
		}
	default:
		return nil, errors.New("not supported filed type")
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
	for i = 0; i < len(tag) && tag[i] != '#'; i++ {
		if !r.noneQuery && tag[i] == '?' {
			break
		}
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
				info.Queries = r.inflateQuery(queryStr)
			}
			tag = tag[i:]
		}
	}

	// check handler if not group field
	if info.handler == "" {
		// assign handler name
		info.handler = utils.UpperFirst(field.Name)
	}

	return info, nil
}

func (r *reflex) anyMethodsFromTag(value string) ([]string, bool) {
	anyMethod := strings.TrimSpace(value)
	methods := strings.Split(anyMethod, ",")
	res := make([]string, 0, len(methods))
	for _, method := range methods {
		method = strings.ToUpper(strings.TrimSpace(method))
		switch method {
		case http.MethodGet,
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodConnect,
			http.MethodOptions,
			http.MethodTrace:
			res = append(res, method)
		}
	}
	if len(res) > 0 {
		return res, true
	}
	return nil, false
}

func (r *reflex) inflateQuery(qs string) []string {
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
func (r *reflex) valueByName(value reflect.Value, name string) interface{} {
	if fieldValue := value.FieldByName(name); !fieldValue.IsNil() {
		return fieldValue.Elem().Interface()
	}
	return nil
}
