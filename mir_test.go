// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package mir_test

import (
	. "github.com/alimy/mir"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fields", func() {
	var (
		tagMirs []*TagMir
		tagMir  *TagMir
		err     error
	)

	Context("check mir custom tag name", func() {
		BeforeEach(func() {
			SetTag("urban")
			if mirs, e := TagMirFrom(&urbanEntry{}); e == nil {
				tagMirs = mirs
				if len(mirs) > 0 {
					tagMir = tagMirs[0]
				}
				err = nil
			} else {
				err = e
			}
			SetTag(DefaultTag)
		})

		It("only one item", func() {
			Expect(tagMirs).Should(HaveLen(1))
		})

		It("tagMir not null", func() {
			Expect(tagMir).ShouldNot(BeNil())
		})

		It("not error", func() {
			Expect(err).Should(BeNil())
		})

		It("had 1 fields", func() {
			Expect(tagMir.Fields).Should(HaveLen(1))
		})

		It("check group", func() {
			Expect(tagMir.Group).To(Equal("v1"))
		})

		It("check fields", func() {
			assertTagFields(tagMir.Fields)
		})
	})
})
