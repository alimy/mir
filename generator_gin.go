package mir

import "errors"

// generatorGin generator for Gin
type generatorGin struct{}

// Name name of generator
func (generatorGin) Name() string {
	return GeneratorGin
}

// Generate generate interface code
func (generatorGin) Generate(entries []*TagMir, opts *GenOpts) error {
	// TODO
	return errors.New("not ready")
}
