// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package core

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/alimy/mir/v3/internal/utils"
)

// EngineInfo Engine information
type EngineInfo struct {
	PkgName     string
	ImportAlias string // import alias name
}

// FieldDescriptor field Descriptor info
type FieldDescriptor struct {
	Imports     map[string]string
	PkgPath     string
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
	Group        string
	Chain        string
	Imports      map[string]string
	PkgPath      string
	PkgName      string
	TypeName     string
	Comment      string // not support now so always empty
	InOuts       []reflect.Type
	Fields       []*FieldDescriptor
	EngineInfo   *EngineInfo
	WatchCtxDone bool
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
	var extSts []reflect.Type
	for _, t := range inOuts {
		if t.PkgPath() == d.PkgPath {
			d.InOuts = append(d.InOuts, t)
		} else {
			extSts = append(extSts, t)
		}
	}
	// to set fields pkg name alias map
	pkgNames := utils.NewStrSet()
	for _, t := range extSts {
		pkgPath := t.PkgPath()
		if pkgPath == "" {
			continue
		}
		pkgs := strings.Split(pkgPath, "/")
		pkgName := pkgs[len(pkgs)-1]
		isAlias := false
		for err := pkgNames.Add(pkgName); err != nil; err = pkgNames.Add(pkgName) {
			isAlias = true
			pkgName = pkgName + "_m"
		}
		if !isAlias {
			pkgName = ""
		}
		d.Imports[pkgPath] = pkgName
	}
	d.setFiledImports()
}

func (d *IfaceDescriptor) setFiledImports() {
	for _, f := range d.Fields {
		f.Imports = d.Imports
	}
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

// IsUseBinding return whether use binding interface
func (d *IfaceDescriptor) IsUseBinding() bool {
	for _, f := range d.Fields {
		if f.In != nil {
			return true
		}
	}
	return false
}

// HttpMethod return http method when f.NotHttpAny() is true
func (f *FieldDescriptor) HttpMethod() string {
	if len(f.HttpMethods) == 1 {
		return f.HttpMethods[0]
	}
	return ""
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
	return utils.QuoteJoin(f.HttpMethods, ",")
}

// InName return In type name
func (f *FieldDescriptor) InName() string {
	if f.In == nil {
		return ""
	}
	pkgPath := f.In.PkgPath()
	if pkgPath == f.PkgPath {
		return f.In.Name()
	}
	return f.aliasPkgName(pkgPath) + "." + f.In.Name()
}

// OutName return Out type name
func (f *FieldDescriptor) OutName() string {
	if f.Out == nil {
		return ""
	}
	pkgPath := f.Out.PkgPath()
	if pkgPath == f.PkgPath {
		return f.Out.Name()
	}
	return f.aliasPkgName(pkgPath) + "." + f.Out.Name()
}

func (f *FieldDescriptor) aliasPkgName(pkgPath string) string {
	if alias := f.Imports[pkgPath]; alias != "" {
		return alias
	}
	pkgs := strings.Split(pkgPath, "/")
	return pkgs[len(pkgs)-1]
}
