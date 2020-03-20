package core

import "testing"

func TestDescriptors(t *testing.T) {
	d := make(Descriptors)
	if err := d.Put(&IfaceDescriptor{
		Group:    "",
		Chain:    "Chain",
		PkgName:  "api",
		TypeName: "site",
		Comment:  "",
		Fields: []*FieldDescriptor{
			{
				Host:       "",
				Path:       "/",
				Queries:    nil,
				HttpMethod: "GET",
				MethodName: "Index",
				Comment:    "",
			},
		},
	}); err != nil {
		t.Error("want nil error but not")
	}
	if _, exist := d.Get(""); !exist {
		t.Error("want exist an item but not")
	}
}
