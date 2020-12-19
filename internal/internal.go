// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package internal

import (
	"context"
	"sync"

	"github.com/alimy/mir/v3/internal/core"

	_ "github.com/alimy/mir/v3/internal/generator"
	_ "github.com/alimy/mir/v3/internal/parser"
)

type mirCtx struct {
	context.Context

	mu            *sync.Mutex
	err           error
	capacity      int
	ifaceChan     chan *core.IfaceDescriptor
	generatorDone chan struct{}
	cancelFunc    context.CancelFunc
}

// Err return cancel error
func (c *mirCtx) Err() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.err
}

// ChanCapacity return ifaceChan's capacity
func (c *mirCtx) Capcity() int {
	return c.capacity
}

// Cancel cancel mir's process logic with an error
func (c *mirCtx) Cancel(err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.err != nil {
		c.err = err
		c.cancelFunc()
	}
}

// ParserDone indicate parser done
func (c *mirCtx) ParserDone() {
	close(c.ifaceChan)
}

// GeneratorDone indicate generator done
func (c *mirCtx) GeneratorDone() {
	close(c.generatorDone)
}

// Wait wait for end of process
func (c *mirCtx) Wait() error {
	select {
	case <-c.Done():
	case <-c.generatorDone:
	}
	return c.Err()
}

// Pipe return source/sink chan *IfaceDescriptor
func (c *mirCtx) Pipe() (<-chan *core.IfaceDescriptor, chan<- *core.IfaceDescriptor) {
	return c.ifaceChan, c.ifaceChan
}

// NewMirCtx return a new mir's context instance
func NewMirCtx(capcity int) core.MirCtx {
	ctx := &mirCtx{
		mu:            &sync.Mutex{},
		capacity:      capcity,
		generatorDone: make(chan struct{}),
		ifaceChan:     make(chan *core.IfaceDescriptor, capcity),
	}
	ctx.Context, ctx.cancelFunc = context.WithCancel(context.Background())
	return ctx
}
