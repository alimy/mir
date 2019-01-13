package mir

import "fmt"

var (
	engine Engine
)

// Setup set engine for register handler
func Setup(e Engine) {
	if engine != nil {
		panic("mir: Setup called twice for engine")
	}
	engine = e
}

// Register use entries's info to register handler to engine
func Register(entries ...interface{}) error {
	if engine == nil {
		return fmt.Errorf("you should need setup a engine instance first then call this function")
	}
	if tagMirs, err := TagMirFrom(entries...); err == nil {
		return engine.Register(tagMirs...)
	} else {
		return err
	}

}
