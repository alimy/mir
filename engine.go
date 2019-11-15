package mir

// Generator list
var (
	GeneratorGin        = "gin"
	GeneratorChi        = "chi"
	GeneratorMux        = "mux"
	GeneratorHttpRouter = "httprouter"
)

// GenOpts generator options
type GenOpts struct {
	Name    string
	OutPath string
}

// Parser parse entries
type Parser interface {
	Name() string
	Parse(entries []interface{}) ([]*TagMir, error)
}

// Generator generate interface code for engine
type Generator interface {
	Name() string
	Generate([]*TagMir, *GenOpts) error
}
