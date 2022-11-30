// Copyright 2022 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package utils

import (
	"errors"
	"net/http"
	"sync"
	"testing"
)

func TestStrSet(t *testing.T) {
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
		s := NewStrSet()
		for _, it := range data.input {
			s.Add(it)
		}

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

func TestOnceSet_Add(t *testing.T) {
	onceSet := newOnceSet()

	if err := onceSet.Add(""); err != nil && err.Error() != "empty item" {
		t.Error("want an error but not")
	}

	if err := onceSet.Add("abc"); err != nil {
		t.Error("want nil error but not")
	}

	if err := onceSet.Add("abc"); err != nil {
		t.Error("want an error but not")
	}
}

func TestOnceSet_Exist(t *testing.T) {
	onceSet := newOnceSet()

	_ = onceSet.Add("abc")
	if exist := onceSet.Exist("abc"); !exist {
		t.Error("want exist an item for 'abc' but not")
	}
}

func TestMuxSet_Add(t *testing.T) {
	muxSet := NewMuxSet()
	if err := muxSet.Add("abc"); err != nil {
		t.Error("want nil error but not")
	}
	if exist := muxSet.Exist("abc"); !exist {
		t.Error("want exist an item for 'abc' but not")
	}
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go tinyTest(wg, t, muxSet)
	go tinyTest(wg, t, muxSet)
	go tinyTest(wg, t, muxSet)
	wg.Wait()
}

func newOnceSet() Set {
	return NewOnceSet(func(it string) error {
		if it == "" {
			return errors.New("empty item")
		}
		return nil
	})
}

func tinyTest(wg *sync.WaitGroup, t *testing.T, set Set) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		if err := set.Add("abc"); err == nil {
			t.Error("want an error but not")
		}
	}
}
