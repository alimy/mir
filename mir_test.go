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

func TestTagMirFrom(t *testing.T) {
	tagMirs, err := TagMirFrom(&site{}, &blog{group: "v2"})
	if err != nil {
		t.Error(err)
	}
	if len(tagMirs) != 2 {
		t.Errorf("want two item but not")
	}
	assertTagMir(t, tagMirs)
}

func assertTagMir(t *testing.T, tagMirs []*TagMir) {
	isCheckedGroupV2 := false
	for _, mir := range tagMirs {
		switch mir.Group {
		case "v1":
			checkGroupV1(t, mir)
		case "v2":
			isCheckedGroupV2 = true
			checkGroupV2(t, mir)
		}
	}

	if !isCheckedGroupV2 {
		t.Errorf("want a v2 group TagMir instance but not")
	}
}

func checkGroupV1(t *testing.T, mir *TagMir) {
	if len(mir.HandlerChain) != 0 {
		t.Errorf("want zero handlerChain but have %d", len(mir.HandlerChain))
	}
	if len(mir.Fields) != 2 {
		t.Errorf("want 2 TagFields but hava %d", len(mir.Fields))
	}
}

func checkGroupV2(t *testing.T, mir *TagMir) {
	if len(mir.HandlerChain) != 2 {
		t.Errorf("want 2 handlerChain but have %d", len(mir.HandlerChain))
	}
	if len(mir.Fields) != 2 {
		t.Errorf("want 2 TagFields but have %d", len(mir.Fields))
	}
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
