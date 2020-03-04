package core

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
