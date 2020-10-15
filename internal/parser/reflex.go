// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package parser

import (
	"reflect"
	"strings"

	"github.com/alimy/mir/v2/core"
	"github.com/alimy/mir/v2/internal/naming"
)

// reflex real parser
type reflex struct {
	engineInfo *core.EngineInfo
	ns         naming.NamingStrategy
	tagName    string
	noneQuery  bool
}

// reflex get Descriptors from parse entries
// Notice: Descriptors may be an empty if no actual item and is not routine safe
func (r *reflex) parse(entries []interface{}) (core.Descriptors, error) {
	var err error
	ds := make(core.Descriptors)
	for _, entry := range entries {
		iface, err := r.ifaceFrom(entry)
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

func (r *reflex) ifaceFrom(entry interface{}) (*core.IfaceDescriptor, error) {
	// used to find tagInfo
	entryType := reflect.TypeOf(entry)
	if entryType == nil {
		return nil, errNilType
	}

	// get real entry type
	isPtr := false
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
		EngineInfo: r.engineInfo,
		TypeName:   entryType.Name(),
		PkgName:    "api", // set default pkg name
		Fields:     make([]*core.FieldDescriptor, 0),
	}
	var groupSetuped, chainSetuped bool
	for i := 0; i < entryType.NumField(); i++ {
		field := entryType.Field(i)
		switch tagInfo, err := r.tagInfoFrom(field); err {
		case nil:
			// group field so just parse group info.group info only have one field
			if tagInfo.isGroup {
				if !groupSetuped {
					groupSetuped = true
					r.inflateGroupInfo(iface, entryValue, tagInfo)
					break
				} else {
					return nil, errMultGroupInfo
				}
			}
			// chain field so just parse chain info only have one field
			if tagInfo.isChain {
				if !chainSetuped {
					iface.Chain = tagInfo.fieldName
					chainSetuped = true
					break
				} else {
					return nil, errMultChainInfo
				}
			}
			iface.Fields = append(iface.Fields, r.fieldFrom(tagInfo))
		case errNotExist:
			// normal field but had no mir tag info so just break to continue process next field
		default:
			return nil, err
		}
	}
	return iface, nil
}

// inflateGroupInfo setup tag group info to TagMir instance
func (r *reflex) inflateGroupInfo(d *core.IfaceDescriptor, v reflect.Value, t *tagInfo) {
	// group field value assign to m.group first or tag group string info assigned
	if group := v.FieldByName(t.fieldName).String(); group != "" {
		d.Group = group
	} else {
		d.Group = t.group
	}
	if d.Group != "" {
		names := strings.Split(d.Group, "/")
		pkgName := r.ns.Naming(names[len(names)-1])
		d.SetPkgName(pkgName)
	}
}

// fieldFrom build tagField from entry and tagInfo
func (r *reflex) fieldFrom(t *tagInfo) *core.FieldDescriptor {
	return &core.FieldDescriptor{
		HttpMethod: t.Method,
		Host:       t.Host,
		Path:       t.Path,
		Queries:    t.Queries,
		MethodName: t.fieldName,
	}
}

func newReflex(info *core.EngineInfo, tagName string, noneQuery bool) *reflex {
	return &reflex{
		engineInfo: info,
		ns:         naming.NewSnakeNamingStrategy(),
		tagName:    tagName,
		noneQuery:  noneQuery,
	}
}
