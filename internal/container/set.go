package container

import (
	"errors"
	"sync"
)

var (
	errExist = errors.New("element had exist")
)

// OnceFunc used after OnceSet add it
type OnceFunc func(string) error

// Set once set
type Set interface {
	Add(string) error
	Exist(string) bool
}

type onceSet struct {
	inSet    map[string]struct{}
	onceFunc OnceFunc
}

type muxSet struct {
	mu    *sync.RWMutex
	inSet map[string]struct{}
}

// Add add a item to set
func (s *onceSet) Add(it string) error {
	if _, exist := s.inSet[it]; !exist {
		err := s.onceFunc(it)
		if err != nil {
			return err
		}
		s.inSet[it] = struct{}{}
	}
	return nil
}

// Exist whether it exist
func (s *onceSet) Exist(it string) bool {
	_, exist := s.inSet[it]
	return exist
}

// Add add a item to set
func (s *muxSet) Add(it string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exist := s.inSet[it]; exist {
		return errExist
	}
	s.inSet[it] = struct{}{}
	return nil
}

// Exist whether it exist
func (s *muxSet) Exist(it string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, exist := s.inSet[it]
	return exist
}

// NewOnceSet return an OnceSet instance
// Warning: OnceSet is not goroutine safe
func NewOnceSet(onceFunc OnceFunc) Set {
	return &onceSet{
		inSet:    make(map[string]struct{}),
		onceFunc: onceFunc,
	}
}

// NewMuxSet return a goroutine safe set
func NewMuxSet() Set {
	return &muxSet{
		mu:    &sync.RWMutex{},
		inSet: make(map[string]struct{}),
	}
}
