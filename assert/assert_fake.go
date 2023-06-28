// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package assert

type fakeTypeAssertor struct{}

func (fakeTypeAssertor) AssertBinding(_obj any) bool {
	return false
}

func (fakeTypeAssertor) AssertRender(_obj any) bool {
	return false
}
