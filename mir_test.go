package mir

import "testing"

func TestRegister(t *testing.T) {
	e := &simpleEngine{pathHandler: make(map[string]handlerFunc)}
	Setup(e)
	Register(&site{})
	handler := e.pathHandler["/index/"]
	ret := handler()
	if ret != "Index" {
		t.Errorf("want Index but actual is %s", ret)
	}
}
