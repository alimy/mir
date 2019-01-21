// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package httprouter_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHttprouter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Httprouter Suite")
}
