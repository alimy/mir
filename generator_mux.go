package mir

import "errors"

// generatorMux generator for Mux
type generatorMux struct{}

// Name name of generator
func (generatorMux) Name() string {
	return GeneratorMux
}

// Generate generate interface code
func (generatorMux) Generate(entries []*TagMir, opts *GenOpts) error {
	// TODO
	return errors.New("not ready")
}
