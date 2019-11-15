package mir

// parserStructTag parse for struct tag
type parserStructTag struct{}

// Name name of parser
func (parserStructTag) Name() string {
	return "parserStructTag"
}

// Parse parse interface define object entries
func (parserStructTag) Parse(entries []interface{}) ([]*TagMir, error) {
	return TagMirFrom(entries...)
}
