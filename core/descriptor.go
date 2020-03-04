package core

// FieldDescriptor field Descriptor info
type FieldDescriptor struct {
	Uri        string
	MethodName string
}

// IfaceDescriptor interface Descriptor info
type IfaceDescriptor struct {
	PkgName     string
	TypeName    string
	IsNeedChain bool
	Fields      []*FieldDescriptor
}

// IfaceDescriptors interface Descriptor map {TypeName:*IfaceDescriptor}
type IfaceDescriptors map[string]*IfaceDescriptor

// Descriptors mir's Descriptor map {group: IfaceDescriptors}
type Descriptors map[string]IfaceDescriptors
