// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package core

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/alimy/mir/v3"
	"github.com/alimy/mir/v3/internal/utils"
)

// EngineInfo Engine information
type EngineInfo struct {
	PkgName     string
	ImportAlias string // import alias name
}

// FieldDescriptor field Descriptor info
type FieldDescriptor struct {
	Host        string
	Path        string
	Queries     []string
	HttpMethods []string
	IsAnyMethod bool
	In          reflect.Type
	Out         reflect.Type
	InOuts      []reflect.Type
	MethodName  string
	Comment     string // not support now so always empty
}

// IfaceDescriptor interface Descriptor info
type IfaceDescriptor struct {
	Group      string
	Chain      string
	PkgName    string
	TypeName   string
	Comment    string // not support now so always empty
	InOuts     []reflect.Type
	Fields     []*FieldDescriptor
	EngineInfo *EngineInfo
}

// IfaceDescriptors interface Descriptor map {TypeName:*IfaceDescriptor}
type IfaceDescriptors map[string]*IfaceDescriptor

// Descriptors mir's Descriptor map {group: IfaceDescriptors}
type Descriptors map[string]IfaceDescriptors

// Put add a IfaceDescriptor
// Notice: if exist same instance by TypeName will override the old one
func (d Descriptors) Put(iface *IfaceDescriptor) error {
	if iface == nil {
		return errors.New("entry is empty")
	}
	key := d.keyFrom(iface.Group)
	ifaces, exist := d[key]
	if !exist {
		ifaces = make(IfaceDescriptors)
		d[key] = ifaces
	}
	if _, exist = ifaces[iface.TypeName]; exist {
		return fmt.Errorf("had exist entry by name: %s", iface.TypeName)
	}
	ifaces[iface.TypeName] = iface
	return nil
}

// Get get a IfaceDescriptors if exists
func (d Descriptors) Get(group string) (IfaceDescriptors, bool) {
	ifaces, exist := d[d.keyFrom(group)]
	return ifaces, exist
}

// Exist whether exist *IfaceDescriptor sub-item
func (d Descriptors) Exist(iface *IfaceDescriptor) bool {
	if iface == nil {
		return false
	}
	key := d.keyFrom(iface.Group)
	ifaces, exist := d[key]
	if !exist {
		return false
	}
	if _, exist = ifaces[iface.TypeName]; !exist {
		return false
	}
	return true
}

// GroupFrom return group name from key
func (d Descriptors) GroupFrom(key string) string {
	return strings.TrimLeft(key, "_")
}

func (d Descriptors) keyFrom(s string) string {
	return "_" + s
}

// SetPkgName set package name
func (d *IfaceDescriptor) SetPkgName(name string) {
	d.PkgName = name
}

// SetInnerInOuts set inner InOuts for defined
func (d *IfaceDescriptor) SetInnerInOuts(inOuts []reflect.Type) {
	d.InOuts = inOuts
}

// AllInOuts return all InOuts from Fileds
func (d *IfaceDescriptor) AllInOuts() []reflect.Type {
	tyns := utils.NewStrSet()
	var inouts []reflect.Type
	for _, f := range d.Fields {
		for _, t := range f.InOuts {
			if !tyns.Exist(t.Name()) {
				inouts = append(inouts, t)
				tyns.Add(t.Name())
			}
		}
	}
	return inouts
}

// NotHttpAny not just http any method
func (f *FieldDescriptor) NotHttpAny() bool {
	return !f.IsAnyMethod && len(f.HttpMethods) == 1
}

// JustHttpAny just http any method
func (f *FieldDescriptor) JustHttpAny() bool {
	return f.IsAnyMethod
}

// AnyHttpMethods return methods in HttpMethods
func (f *FieldDescriptor) AnyHttpMethods() []string {
	return f.HttpMethods
}

// HttpMethodArgs return http method as argument like "POST","GET","HEAD"
func (f *FieldDescriptor) HttpMethodArgs() string {
	httpMthods := mir.HttpMethods
	return utils.QuoteJoin(httpMthods, ",")
}

// InName return In type name
func (f *FieldDescriptor) InName() string {
	if f.In != nil {
		return f.In.Name()
	}
	return ""
}

// OutName return Out type name
func (f *FieldDescriptor) OutName() string {
	if f.Out != nil {
		return f.OutName()
	}
	return ""
}
