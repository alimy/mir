package mir

import "errors"

// generatorHttpRouter generator for HttpRouter
type generatorHttpRouter struct{}

// Name name of generator
func (generatorHttpRouter) Name() string {
	return GeneratorHttpRouter
}

// Generate generate interface code
func (generatorHttpRouter) Generate(entries []*TagMir, opts *GenOpts) error {
	// TODO
	return errors.New("not ready")
}
