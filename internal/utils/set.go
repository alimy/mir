// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package utils

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
	List() []string
}

type onceSet struct {
	inSet    map[string]struct{}
	onceFunc OnceFunc
}

type muxSet struct {
	mu    *sync.RWMutex
	inSet map[string]struct{}
}

type strSet map[string]struct{}

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

// List return all items in sets
func (s *onceSet) List() []string {
	res := make([]string, 0, len(s.inSet))
	for item := range s.inSet {
		res = append(res, item)
	}
	return res
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

// List return all items in sets
func (s *muxSet) List() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	res := make([]string, 0, len(s.inSet))
	for item := range s.inSet {
		res = append(res, item)
	}
	return res
}

// Add add a item to set
func (s strSet) Add(item string) error {
	if _, exist := s[item]; exist {
		return errExist
	}
	s[item] = struct{}{}
	return nil
}

// Exist whether it exist
func (s strSet) Exist(item string) bool {
	_, exist := s[item]
	return exist
}

// List return all items in sets
func (s strSet) List() []string {
	methods := make([]string, 0, len(s))
	for m := range s {
		methods = append(methods, m)
	}
	return methods
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

// NewStrSet return a str sets
func NewStrSet() Set {
	return make(strSet)
}
