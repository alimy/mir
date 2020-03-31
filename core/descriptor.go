package core

import (
	"errors"
	"fmt"
	"strings"
)

// FieldDescriptor field Descriptor info
type FieldDescriptor struct {
	Host       string
	Path       string
	Queries    []string
	HttpMethod string
	MethodName string
	Comment    string // not support now so always empty
}

// IfaceDescriptor interface Descriptor info
type IfaceDescriptor struct {
	Group    string
	Chain    string
	PkgName  string
	TypeName string
	Comment  string // not support now so always empty
	Fields   []*FieldDescriptor
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
