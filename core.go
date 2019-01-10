package mir

// Group indicator a default group for handler to register to server engine
type Group string

// Get indicator a GET method handler used placeholder register info in struct tag
type Get struct{}

// Put indicator a PUT method handler used placeholder register info in struct tag
type Put struct{}

// Post indicator a POST method handler used placeholder register info in struct tag
type Post struct{}

// Delete indicator a DELETE method handler used placeholder register info in struct tag
type Delete struct{}

// Head indicator a HEAD method handler used placeholder register info in struct tag
type Head struct{}

// Patch indicator a PATCH method handler used placeholder register info in struct tag
type Patch struct{}

// Options indicator a OPTIONS method handler used placeholder register info in struct tag
type Options struct{}

// Any indicator a Any method handler used placeholder register info in struct tag.
// This is mean register handler that all http.Method* include(GET, PUT, POST, DELETE,
// HEAD, PATCH, OPTIONS)
type Any struct{}

type Engine interface {
	Register(...interface{}) error
}
