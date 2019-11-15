package mir

import "errors"

// generatorChi generator for go-chi
type generatorChi struct{}

// Name name of generator
func (generatorChi) Name() string {
	return GeneratorChi
}

// Generate generate interface code
func (generatorChi) Generate(entries []*TagMir, opts *GenOpts) error {
	// TODO
	return errors.New("not ready")
}
