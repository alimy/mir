// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package mir

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

const (
	// DefaultTag indicate default mir's struct tag name
	DefaultTag = "mir"
)

var (
	// tagName indicate mir's struct tag string name
	tagName = DefaultTag
)

var (
	tagNotExist      tagError = "mir struct tag filed not exist"
	tagNoPathInfo    tagError = "mir struct tag not contains path info"
	tagNotValideType tagError = "not valide type, just struct and struct ptr is avalibale"
	tagMultGroupInfo tagError = "multiple group info in struct defined"
	tagMultChainInfo tagError = "multiple chain info in struct defined"
)

// tagError indicate error information
type tagError string

func (e tagError) Error() string {
	return string(e)
}

// tagBase indicate mir tag common information in struct tag string exclude handler
type tagBase struct {
	Method  string   // Method indicate method information in struct tag string
	Host    string   // Host indicate host information in struct tag string
	Path    string   // Path indicate path information in struct tag string
	Queries []string // Queries indicate path information in struct tag string
}

// tagInfo indicate mir tag information in struct tag string
type tagInfo struct {
	tagBase
	isGroup   bool   // indicate whether a group field
	group     string // indicate group information in struct tag string
	groupName string // indicate group field name
	chainName string // indicate chain field name
	handler   string // indicate handler information in struct tag string
}

// TagField indicate mir tag info used to register route info to engine
type TagField struct {
	tagBase
	Handler interface{} // Handler indicate handler method
}

// TagMir contains TagFields by group
type TagMir struct {
	Group  string
	Chain  Chain
	Fields []*TagField
}

// tagFieldsGroup indicate group-tagFields map
type tagFieldsGroup map[string]*TagMir

// TagMirFrom build TagMir items from entries slice
func TagMirFrom(entries ...interface{}) ([]*TagMir, error) {
	mergedTagMirs := make(tagFieldsGroup)
	for _, entry := range entries {
		if tagFields, err := tagMirFrom(entry); err == nil {
			// no actual field so just continue
			if len(tagFields.Fields) == 0 {
				continue
			}
			// merge tagFields by group
			tagMir, exist := mergedTagMirs[tagFields.Group]
			if !exist {
				tagMir = &TagMir{
					Group:  tagFields.Group,
					Fields: make([]*TagField, 0, len(tagFields.Fields)),
				}
				mergedTagMirs[tagFields.Group] = tagMir
			}
			// Notice: override tagMir.Chain if tagFields.Chain not nil
			if tagFields.Chain != nil {
				tagMir.Chain = tagFields.Chain
			}
			tagMir.Fields = append(tagMir.Fields, tagFields.Fields...)
		} else {
			return nil, err
		}
	}

	// build result
	tagMirSlice := make([]*TagMir, 0, len(mergedTagMirs))
	for _, item := range mergedTagMirs {
		tagMirSlice = append(tagMirSlice, item)
	}
	return tagMirSlice, nil
}

// tagMirFrom build tagMir items from a entry
func tagMirFrom(entry interface{}) (*TagMir, error) {
	// used to find tagInfo
	entryType := reflect.TypeOf(entry)
	isPtr := false

	// get real entry type
	if entryType.Kind() == reflect.Ptr {
		isPtr = true
		entryType = entryType.Elem()
	}

	// entry must struct type
	if entryType.Kind() != reflect.Struct {
		return nil, tagNotValideType
	}

	// used to find method from T and lookup struct tag string of mir tag information
	var entryValue, entryPtrValue reflect.Value
	if isPtr {
		entryPtrValue = reflect.ValueOf(entry)
		entryValue = entryPtrValue.Elem()
	} else {
		entryValue = reflect.ValueOf(entry)
		entryPtrValue = entryValue.Addr()
	}

	// get tagMir from entryType and entryPtrType
	tagMir := &TagMir{Fields: make([]*TagField, 0)}
	groupSetuped := false
	chainSetuped := false
	for i := 0; i < entryType.NumField(); i++ {
		field := entryType.Field(i)
		switch tagInfo, err := tagInfoFrom(field); err {
		case nil:
			// group field so just parse group info.group info only have one field
			if tagInfo.isGroup {
				if !groupSetuped {
					groupSetuped = true
					inflateGroupInfo(tagMir, entryValue, tagInfo)
					break
				} else {
					return nil, tagMultGroupInfo
				}
			}
			// chain field so just parse chain info only have one field
			if tagInfo.chainName != "" {
				if !chainSetuped {
					if value := entryValue.FieldByName(tagInfo.chainName); !value.IsNil() {
						tagMir.Chain = entryValue.FieldByName(tagInfo.chainName).Elem().Interface()
						chainSetuped = true
					}
					break
				} else {
					return nil, tagMultChainInfo
				}
			}
			// method field build tagField first
			if tagField, err := tagFieldFrom(entryValue, entryPtrValue, tagInfo); err == nil {
				tagMir.Fields = append(tagMir.Fields, tagField)
			} else {
				return nil, err
			}
		case tagNotExist:
			// normal field but had no mir tag info so just break to continue process next field
		default:
			return nil, err
		}
	}
	return tagMir, nil
}

// tagFieldFrom build tagField from entry and tagInfo
func tagFieldFrom(v reflect.Value, ptrV reflect.Value, t *tagInfo) (*TagField, error) {
	if m, err := methodByName(v, ptrV, t.handler); err == nil {
		return &TagField{tagBase: t.tagBase, Handler: m}, nil
	} else {
		return nil, err
	}
}

// tagInfoFrom build tagInfo from field
func tagInfoFrom(field reflect.StructField) (*tagInfo, error) {
	info := &tagInfo{}

	// lookup mir tag info from struct field
	tag, exist := field.Tag.Lookup(tagName)
	if !exist {
		return nil, tagNotExist
	}

	// Skip leading space.
	i := 0
	for i < len(tag) && tag[i] == ' ' {
		i++
	}
	tag = tag[i:]

	// group info or method info or chain info
	switch field.Type.Name() {
	case "Chain":
		info.chainName = field.Name
		return info, nil
	case "Group":
		info.isGroup = true
		info.group = tag
		info.groupName = field.Name
		return info, nil
	case "Get":
		info.Method = MethodGet
	case "Put":
		info.Method = MethodPut
	case "Post":
		info.Method = MethodPost
	case "Delete":
		info.Method = MethodDelete
	case "Head":
		info.Method = MethodHead
	case "Options":
		info.Method = MethodOptions
	case "Patch":
		info.Method = MethodPatch
	case "Trace":
		info.Method = MethodTrace
	case "Connect":
		info.Method = MethodConnect
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
		return nil, tagNoPathInfo
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
			info.handler = tag[1:i]
			tag = tag[i:]
		case '?':
			i := 1
			for i < len(tag) && tag[i] != '#' {
				i++
			}
			queryStr := tag[1:i]
			if queryStr != "" {
				info.Queries = strings.Split(queryStr, "&")
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

// inflateGroupInfo setup tag group info to TagMir instance
func inflateGroupInfo(m *TagMir, v reflect.Value, t *tagInfo) {
	// group field value assign to m.group first or tag group string info assigned
	if group := v.FieldByName(t.groupName).String(); group != "" {
		m.Group = group
	} else {
		m.Group = t.group
	}
}

// methodByName return a method interface from value and method name
func methodByName(value reflect.Value, ptrValue reflect.Value, name string) (m interface{}, err error) {
	if v := value.MethodByName(name); v.IsValid() {
		m = v.Interface()
	} else if v := ptrValue.MethodByName(name); v.IsValid() {
		m = v.Interface()
	} else {
		typeName := value.Type().Name()
		err = fmt.Errorf("not found method %s in struct type %s or *%s", name, typeName, typeName)
	}
	return
}
