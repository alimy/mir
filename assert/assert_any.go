// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package assert

// anyTypeAssertor a common type assert for type T
type anyTypeAssertor[T any] struct{}

func (anyTypeAssertor[T]) AssertBinding(obj any) bool {
	_, ok := obj.(Binding[T])
	return ok
}

func (anyTypeAssertor[T]) AssertRender(obj any) bool {
	_, ok := obj.(Render[T])
	return ok
}
