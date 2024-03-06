// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package core

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/alimy/mir/v4/internal/utils"
)

var (
	VerInfo = &VersionInfo{
		MirVer: "v4.2.0",
	}
)

// EngineInfo Engine information
type EngineInfo struct {
	PkgName     string
	ImportAlias string // import alias name
}

// VersionInfo mir version information
type VersionInfo struct {
	MirVer string
}

// FieldDescriptor field Descriptor info
type FieldDescriptor struct {
	Imports             map[string]string
	PkgPath             string
	Host                string
	Path                string
	Queries             []string
	HttpMethods         []string
	IsAnyMethod         bool
	IsFieldChain        bool
	IsUseContext        bool
	IsUseRequestContext bool
	IsBindIn            bool
	IsRenderOut         bool
	BindingName         string
	RenderName          string
	In                  reflect.Type
	Out                 reflect.Type
	InOuts              []reflect.Type
	MethodName          string
	Comment             string // not support now so always empty
}

// IfaceDescriptor interface Descriptor info
type IfaceDescriptor struct {
	Group                string
	Chain                string
	Imports              map[string]string
	PkgPath              string
	PkgName              string
	TypeName             string
	Comment              string // not support now so always empty
	InOuts               []reflect.Type
	Fields               []*FieldDescriptor
	EngineInfo           *EngineInfo
	VerInfo              *VersionInfo
	WatchCtxDone         bool
	UseRequestCtx        bool
	DeclareCoreInterface bool // whether need to declare core interface, default is false
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

// SortedIfaces return sorted Iface slice
func (d IfaceDescriptors) SortedIfaces() []*IfaceDescriptor {
	keys := make([]string, 0, len(d))
	ifaces := make([]*IfaceDescriptor, 0, len(d))
	for key := range d {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		ifaces = append(ifaces, d[key])
	}
	return ifaces
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

// SetDeclareCoreInterface set declare core interface value
func (d *IfaceDescriptor) SetDeclareCoreInterface(isNeed bool) {
	d.DeclareCoreInterface = isNeed
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
	pkgNames := make(map[string]string)
	for _, t := range extSts {
		pkgPath := t.PkgPath()
		if pkgPath == "" {
			continue
		}
		// had import so no need process
		if _, exist := d.Imports[pkgPath]; exist {
			continue
		}
		// process alias if needed
		pkgs := strings.Split(pkgPath, "/")
		pkgName := pkgs[len(pkgs)-1]
		if pkg, exist := pkgNames[pkgName]; !exist {
			pkgNames[pkgName] = pkgPath
			d.Imports[pkgPath] = ""
		} else {
			for exist && pkg != pkgPath {
				pkgName = pkgName + "_m"
				pkg, exist = pkgNames[pkgName]
			}
			pkgNames[pkgName] = pkgPath
			d.Imports[pkgPath] = pkgName
		}
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

// ChainFields return field chains
func (d *IfaceDescriptor) ChainFields() (fields []*FieldDescriptor) {
	for _, f := range d.Fields {
		if f.IsFieldChain {
			fields = append(fields, f)
		}
	}
	return
}

// IsUseFieldChain whether use field chain
func (d *IfaceDescriptor) IsUseFieldChain() bool {
	for _, f := range d.Fields {
		if f.IsFieldChain {
			return true
		}
	}
	return false
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

func (d *IfaceDescriptor) IsUseRequestContext() bool {
	if !d.UseRequestCtx {
		return false
	}
	for _, f := range d.Fields {
		if f.IsUseRequestContext {
			return true
		}
	}
	return false
}

// BindingFields return Binding's fields
func (d *IfaceDescriptor) BindingFields() (fields []*FieldDescriptor) {
	for _, f := range d.Fields {
		if f.In != nil {
			fields = append(fields, f)
		}
	}
	return
}

func (d *IfaceDescriptor) IsUseNamedBinding() bool {
	for _, f := range d.Fields {
		if f.IsUseNamedBinding() {
			return true
		}
	}
	return false
}

func (d *IfaceDescriptor) IsUseNamedRender() bool {
	for _, f := range d.Fields {
		if f.IsUseNamedRender() {
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

// JustUseContext whether just use context
func (f *FieldDescriptor) JustUseContext() bool {
	return f.IsUseContext && len(f.InOuts) == 0
}

// OrInOut in or out
func (f *FieldDescriptor) OrInOut() bool {
	return f.In != nil || f.Out != nil
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

func (f *FieldDescriptor) IsUseNamedBinding() bool {
	return f.In != nil && len(f.BindingName) > 0
}

func (f *FieldDescriptor) IsUseNamedRender() bool {
	return f.Out != nil && len(f.RenderName) > 0
}

func (f *FieldDescriptor) aliasPkgName(pkgPath string) string {
	if alias := f.Imports[pkgPath]; alias != "" {
		return alias
	}
	pkgs := strings.Split(pkgPath, "/")
	return pkgs[len(pkgs)-1]
}
