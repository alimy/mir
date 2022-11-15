package utils

import (
	"net/http"
	"testing"
)

func TestMethodSet(t *testing.T) {
	for _, data := range []struct {
		input  []string
		expect []string
		exist  string
	}{
		{
			input:  []string{http.MethodGet},
			expect: []string{http.MethodGet},
			exist:  http.MethodGet,
		},
		{
			input:  []string{"others"},
			expect: []string{"others"},
			exist:  "others",
		},
	} {
		s := Set{}
		s.Add(data.input...)

		if !s.Exist(data.exist) {
			t.Errorf("want exist %s but not", data.exist)
		}

		list := s.List()
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
