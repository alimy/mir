package mir

import "testing"

func TestRegisterDefault(t *testing.T) {
	e := &simpleEngine{pathHandler: make(map[string]handlerFunc)}
	SetDefault(e)
	if err := RegisterDefault(&site{}); err != nil {
		t.Error(err)
		return
	}
	assertSimpleEngine(t, e)
}

func TestRegister(t *testing.T) {
	e := &simpleEngine{pathHandler: make(map[string]handlerFunc)}
	if err := Register(e, &site{}); err != nil {
		t.Error(err)
		return
	}
	assertSimpleEngine(t, e)
}

func assertSimpleEngine(t *testing.T, e *simpleEngine) {
	handler := e.pathHandler["/index/"]
	if handler != nil {
		ret := handler()
		if ret != "Index" {
			t.Errorf("want Index but actual is %s", ret)
		}
	} else {
		t.Errorf("not register success")
	}
}
