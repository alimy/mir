package core

import (
	"strings"
	"unicode"
)

// FieldDescriptor field Descriptor info
type FieldDescriptor struct {
	Uri        string
	HttpMethod string
	Comment    string
	MethodName string
}

// IfaceDescriptor interface Descriptor info
type IfaceDescriptor struct {
	Group       string
	PkgName     string
	TypeName    string
	Comment     string
	IsNeedChain bool
	Fields      []*FieldDescriptor
}

// IfaceDescriptors interface Descriptor map {TypeName:*IfaceDescriptor}
type IfaceDescriptors map[string]*IfaceDescriptor

// Descriptors mir's Descriptor map {group: IfaceDescriptors}
type Descriptors map[string]IfaceDescriptors

// Add add a IfaceDescriptor
// Notice: if exist same instance by TypeName will override the old one
func (d Descriptors) Add(iface *IfaceDescriptor) {
	if iface == nil {
		return
	}
	key := d.keyFrom(iface.Group)
	ifaces, exist := d[key]
	if !exist {
		ifaces = make(IfaceDescriptors, 2)
		d[key] = ifaces
	}
	ifaces[iface.TypeName] = iface
}

// GroupFrom return group name from key
func (d Descriptors) GroupFrom(key string) string {
	return strings.TrimLeft(key, "_")
}

func (d Descriptors) keyFrom(s string) string {
	return "_" + s
}

// SnakeFileName return snake file name for interface define file
func (d *IfaceDescriptor) SnakeFileName() string {
	var b strings.Builder
	notFirst := false
	b.Grow(len(d.TypeName) + 3)
	for _, r := range d.TypeName {
		if unicode.IsUpper(r) {
			if notFirst {
				b.WriteRune('_')
			}
			b.WriteRune(unicode.ToLower(r))
		} else {
			b.WriteRune(r)
		}
		notFirst = true
	}
	b.WriteString(".go")
	return b.String()
}
