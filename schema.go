package mir

type (
	Interface interface {
		mustEmbedSchema()
	}
	// Schema is the default implementation for the restful Interface.
	// It can be embedded in end-user schemas as follows:
	//
	//	type T struct {
	//		mir.Schema
	//	}
	//
	Schema struct{}
)

func (Schema) mustEmbedSchema() {}
