package mir

import (
	"net/http"
	"testing"
)

func TestMethodSet_Add_List(t *testing.T) {
	for _, data := range []struct {
		input  []string
		expect []string
	}{
		{
			input:  []string{http.MethodGet},
			expect: []string{http.MethodGet},
		},
		{
			input:  []string{"others"},
			expect: []string{},
		},
		{
			input:  HttpMethods,
			expect: HttpMethods,
		},
		{
			input: []string{
				http.MethodGet,
				http.MethodHead,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
				http.MethodConnect,
				http.MethodOptions,
				http.MethodTrace,
				"others",
				"notbeincludes",
			},
			expect: HttpMethods,
		},
	} {
		ms := MethodSet{}
		ms.Add(data.input...)
		list := ms.List()
		if len(list) != len(data.expect) {
			t.Errorf("want list length=%d but got %d", len(data.expect), len(list))
		}

	Top:
		for _, lv := range list {
			for _, ev := range data.expect {
				if lv == ev {
					continue Top
				}
			}
			t.Errorf("want list %v but got %v", data.expect, list)
		}
	}
}
