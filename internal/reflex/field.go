// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package reflex

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/alimy/mir/v5"
	"github.com/alimy/mir/v5/assert"
	"github.com/alimy/mir/v5/internal/utils"
)

var (
	// error list
	errNoop           tagError = "noop"
	errNilType        tagError = "nil type is not valide"
	errNotExist       tagError = "mir struct tag filed not exist"
	errNoPathInfo     tagError = "mir struct tag not contains path info"
	errNotValideType  tagError = "not valide type, just struct and struct ptr is avalibale"
	errMultGroupInfo  tagError = "multiple group info in struct defined"
	errMultChainInfo  tagError = "multiple chain info in struct defined"
	errMultSchemaInfo tagError = "multiple schema info in struct defined"
)

const (
	mirPkgName = "github.com/alimy/mir/v5"
)

// tagError indicate error information
type tagError string

// Error error message string
func (e tagError) Error() string {
	return string(e)
}

// tagInfo indicate mir tag information in struct tag string
type tagInfo struct {
	isAnyMethod  bool                // isAnyMethod indicate whether method is Any
	isFieldChain bool                // isFieldChain indicate whether method is need field chain
	isUseContext bool                // isUseContext indicate whether method is just use Context
	methods      utils.HttpMethodSet // method indicate methods information in struct tag string
	host         string              // host indicate host information in struct tag string
	path         string              // path indicate path information in struct tag string
	queries      []string            // queries indicate path information in struct tag string
	isSchema     bool                // indicate whether a schema field
	isGroup      bool                // indicate whether a group field
	isChain      bool                // indicate whether a chain field
	group        string              // indicate group information in struct tag string
	chainFunc    string              // indicate chain function information in struct tag string
	handler      string              // indicate handler information in struct tag string
	fieldName    string              // indicate field name
	schemaChain  string
	isBindIn     bool
	isRenderOut  bool
	bindingName  string
	renderName   string
	in           reflect.Type
	out          reflect.Type
	inOuts       []reflect.Type
	comment      string // indicate comment info (not support now)
}

// tagInfoFrom build tagInfo from field
func (r *reflex) tagInfoFrom(field reflect.StructField, pkgPath string) (*tagInfo, error) {
	info := &tagInfo{
		methods: make(utils.HttpMethodSet, 1),
	}

	// lookup mir tag info from struct field
	tag, exist := field.Tag.Lookup(r.tagName)
	if !exist {
		return nil, errNotExist
	}
	// lookup binding/render tag info from struct field
	info.bindingName, _ = field.Tag.Lookup(r.bindingTagName)
	info.renderName, _ = field.Tag.Lookup(r.renderTagName)

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
		if field.Type.PkgPath() != mirPkgName {
			return nil, errors.New("parser[1] not supported filed type")
		}
		switch field.Type.Name() {
		case "Chain":
			info.isChain = true
			return info, nil
		case "Group":
			info.isGroup = true
			info.group = tag
			return info, nil
		case "Schema":
			tagVals := strings.Split(tag, ",")
			switch len(tagVals) {
			case 0:
				// just pass this case
				return nil, errNoop
			case 1:
				info.group = tagVals[0]
			case 2:
				fallthrough
			default:
				info.group = tagVals[0]
				if strings.TrimSpace(tagVals[1]) == "chain" {
					info.schemaChain = "Chain"
				}
			}
			info.isSchema = true
			info.group = strings.TrimSpace(info.group)
			return info, nil
		default:
			return nil, errors.New("parser[2] not supported filed type")
		}
	case reflect.Struct:
		if field.Type.Name() != "Schema" || field.Type.PkgPath() != mirPkgName {
			return nil, errors.New("parser[3] not supported filed type")
		}
		tagVals := strings.Split(tag, ",")
		switch len(tagVals) {
		case 0:
			// just pass this case
			return nil, errNoop
		case 1:
			info.group = tagVals[0]
		case 2:
			fallthrough
		default:
			info.group = tagVals[0]
			if strings.TrimSpace(tagVals[1]) == "chain" {
				info.schemaChain = "Chain"
			}
		}
		info.group = strings.TrimSpace(info.group)
		return info, nil
	case reflect.Func:
		ft := field.Type
		numIn := ft.NumIn()
		numOut := ft.NumOut()
		if numOut > 1 {
			return nil, errors.New("func field just need one most return value")
		}
		if numIn > 0 {
			// request type in latest in argument if declared
			it := ft.In(numIn - 1)
			if it.Kind() == reflect.Struct {
				cts, err := CheckStruct(it, pkgPath)
				if err != nil {
					return nil, err
				}
				info.in = it
				if it.PkgPath() != pkgPath {
					info.isBindIn = assert.AssertBinding(reflect.New(it).Interface())
				}
				info.inOuts = append(info.inOuts, cts...)

				// minus numIn to ignore latest in argument that had processed
				numIn--
			}

			// process other in argument
			for i := numIn - 1; i >= 0; i-- {
				it = ft.In(i)
				if it.PkgPath() != mirPkgName {
					continue
				}
				switch it.Name() {
				case "Get":
					info.methods.Add(mir.MethodGet)
				case "Put":
					info.methods.Add(mir.MethodPut)
				case "Post":
					info.methods.Add(mir.MethodPost)
				case "Delete":
					info.methods.Add(mir.MethodDelete)
				case "Head":
					info.methods.Add(mir.MethodHead)
				case "Options":
					info.methods.Add(mir.MethodOptions)
				case "Patch":
					info.methods.Add(mir.MethodPatch)
				case "Trace":
					info.methods.Add(mir.MethodTrace)
				case "Connect":
					info.methods.Add(mir.MethodConnect)
				case "Any":
					info.isAnyMethod = true
					info.methods.Add(mir.HttpMethods...)
				case "Chain":
					info.isFieldChain = true
				case "Context":
					info.isUseContext = true
				}
			}
		}
		// process special case for not set methods
		if len(info.methods) == 0 {
			info.isAnyMethod = true
			info.methods.Add(mir.HttpMethods...)
		}
		if numOut == 1 {
			ot := ft.Out(i)
			if ot.Kind() != reflect.Struct {
				return nil, errors.New("func field must return value is need struct type")
			}
			cts, err := CheckStruct(ot, pkgPath)
			if err != nil {
				return nil, err
			}
			info.out = ot
			if ot.PkgPath() != pkgPath {
				info.isRenderOut = assert.AssertRender(reflect.New(ot).Interface())
			}
			info.inOuts = append(info.inOuts, cts...)
		}
	default:
		return nil, errors.New("parser[4] not supported filed type")
	}

	// host info
	if len(tag) > 2 && tag[0] == '/' && tag[1] == '/' {
		i := 2
		for i < len(tag) && tag[i] != '/' {
			i++
		}
		info.host = tag[2:i]
		tag = tag[i:]
	}

	// path info. must have path info if not a group field
	if len(tag) == 0 && !info.isGroup {
		return nil, fmt.Errorf("invalid tag info by pkg_path:%s entry_field:%s err: %w", pkgPath, info.fieldName, errNoPathInfo)
	}
	for i = 0; i < len(tag) && tag[i] != '#'; i++ {
		if !r.noneQuery && tag[i] == '?' {
			break
		}
	}
	info.path = tag[0:i]
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
				info.queries = r.inflateQuery(queryStr)
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
func (r *reflex) valueByName(value reflect.Value, name string) any {
	if fieldValue := value.FieldByName(name); !fieldValue.IsNil() {
		return fieldValue.Elem().Interface()
	}
	return nil
}
