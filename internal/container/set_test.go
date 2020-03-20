package container

import (
	"errors"
	"sync"
	"testing"
)

func newOnceSet() Set {
	return NewOnceSet(func(it string) error {
		if it == "" {
			return errors.New("empty item")
		}
		return nil
	})
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

func tinyTest(wg *sync.WaitGroup, t *testing.T, set Set) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		if err := set.Add("abc"); err == nil {
			t.Error("want an error but not")
		}
	}
}
