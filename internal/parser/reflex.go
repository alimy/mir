// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package parser

import (
	"reflect"

	"github.com/alimy/mir/v2/core"
)

// reflex get Descriptors from parse entries
// Notice: Descriptors may be an empty if no actual item
func reflex(entries []interface{}) (core.Descriptors, error) {
	var err error
	ds := make(core.Descriptors)
	for _, entry := range entries {
		iface, err := ifaceFrom(entry)
		if err != nil {
			break
		}
		// no actual fields so just continue
		if len(iface.Fields) == 0 {
			continue
		}
		if err = ds.Put(iface); err != nil {
			break
		}
	}
	return ds, err
}

func ifaceFrom(entry interface{}) (*core.IfaceDescriptor, error) {
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
		return nil, errNotValideType
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

	// get IfaceDescriptor from entryType and entryPtrType
	iface := &core.IfaceDescriptor{
		TypeName: entryType.Name(),
		Fields:   make([]*core.FieldDescriptor, 0),
	}
	var groupSetuped, chainSetuped bool
	for i := 0; i < entryType.NumField(); i++ {
		field := entryType.Field(i)
		switch tagInfo, err := tagInfoFrom(field); err {
		case nil:
			// group field so just parse group info.group info only have one field
			if tagInfo.isGroup {
				if !groupSetuped {
					groupSetuped = true
					inflateGroupInfo(iface, entryValue, tagInfo)
				} else {
					return nil, errMultGroupInfo
				}
			}
			// chain field so just parse chain info only have one field
			if tagInfo.isChain {
				if !chainSetuped {
					iface.IsNeedChain = true
					chainSetuped = true
				} else {
					return nil, errMultChainInfo
				}
			}
			iface.Fields = append(iface.Fields, fieldFrom(tagInfo))
		case errNotExist:
			// normal field but had no mir tag info so just break to continue process next field
		default:
			return nil, err
		}
	}
	return iface, nil
}

// inflateGroupInfo setup tag group info to TagMir instance
func inflateGroupInfo(d *core.IfaceDescriptor, v reflect.Value, t *tagInfo) {
	// group field value assign to m.group first or tag group string info assigned
	if group := v.FieldByName(t.fieldName).String(); group != "" {
		d.Group = group
	} else {
		d.Group = t.group
	}
}

// fieldFrom build tagField from entry and tagInfo
func fieldFrom(t *tagInfo) *core.FieldDescriptor {
	return &core.FieldDescriptor{
		HttpMethod: t.Method,
		Host:       t.Host,
		Path:       t.Path,
		Queries:    t.Queries,
		MethodName: t.fieldName,
	}
}
