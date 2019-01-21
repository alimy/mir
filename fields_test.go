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
		tagMirs     []*TagMir
		tagMir      *TagMir
		err         error
		commonPaths []string
	)

	BeforeEach(func() {
		commonPaths = []string{
			"/get/",
			"/put/",
			"/post/",
			"/delete/",
			"/head/",
			"/patch/",
			"/trace/",
			"/connect/",
			"/options/",
			"/any/",
			"/alias/",
		}
	})

	Context("mux style URN tag string", func() {
		BeforeEach(func() {
			tagMirs, err = TagMirFrom(&muxEntry{Chain: mirChains()})
			if err == nil && len(tagMirs) > 0 {
				tagMir = tagMirs[0]
			}
			commonPaths = append(commonPaths, "/query/", "/full/{other}/{id:[0-9]+}")
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

		It("had 13 fields", func() {
			Expect(tagMir.Fields).Should(HaveLen(13))
		})

		It("check group", func() {
			Expect(tagMir.Group).To(Equal("v1"))
		})

		It("check chain", func() {
			chains, ok := tagMir.Chain.(chains)
			Expect(ok).Should(BeTrue())
			Expect(chains).Should(HaveLen(2))
		})

		It("check fields", func() {
			fields := assertTagFields(tagMir.Fields)
			for _, path := range commonPaths {
				Expect(fields).Should(HaveKey(path))
			}
		})
	})

	Context("gin style URN tag string", func() {
		BeforeEach(func() {
			tagMirs, err = TagMirFrom(&ginEntry{Chain: mirChains()})
			if err == nil && len(tagMirs) > 0 {
				tagMir = tagMirs[0]
			}
			commonPaths = append(commonPaths, "/full/:other/:name")
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

		It("had 12 fields", func() {
			Expect(tagMir.Fields).Should(HaveLen(12))
		})

		It("check group", func() {
			Expect(tagMir.Group).To(Equal("v1"))
		})

		It("check chain", func() {
			chains, ok := tagMir.Chain.(chains)
			Expect(ok).Should(BeTrue())
			Expect(chains).Should(HaveLen(2))
		})

		It("check fields", func() {
			fields := assertTagFields(tagMir.Fields)
			for _, path := range commonPaths {
				Expect(fields).Should(HaveKey(path))
			}
		})
	})

	Context("iris style URN tag string", func() {
		BeforeEach(func() {
			tagMirs, err = TagMirFrom(&irisEntry{Chain: mirChains()})
			if err == nil && len(tagMirs) > 0 {
				tagMir = tagMirs[0]
			}
			commonPaths = append(commonPaths, "/full/{other:string}/{name:string range(1,200) else 400}")
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

		It("had 12 fields", func() {
			Expect(tagMir.Fields).Should(HaveLen(12))
		})

		It("check group", func() {
			Expect(tagMir.Group).To(Equal("v1"))
		})

		It("check chain", func() {
			chains, ok := tagMir.Chain.(chains)
			Expect(ok).Should(BeTrue())
			Expect(chains).Should(HaveLen(2))
		})

		It("check fields", func() {
			fields := assertTagFields(tagMir.Fields)
			for _, path := range commonPaths {
				Expect(fields).Should(HaveKey(path))
			}
		})
	})

	Context("tagMirs from 3 entries", func() {
		BeforeEach(func() {
			entries := []interface{}{
				&muxEntry{Group: "v2"},
				&ginEntry{Chain: mirChains()},
				&irisEntry{},
			}
			tagMirs, err = TagMirFrom(entries...)
		})

		It("not error", func() {
			Expect(err).Should(BeNil())
		})

		It("want 2 item", func() {
			Expect(tagMirs).Should(HaveLen(2))
		})

		It("check group", func() {
			haveV2Group := false
			for _, item := range tagMirs {
				if item.Group == "v2" {
					haveV2Group = true
					Expect(item.Fields).Should(HaveLen(13))
				} else if item.Group == "v1" {
					Expect(item.Fields).Should(HaveLen(24))
				}
			}
			Expect(haveV2Group).Should(BeTrue())
		})
	})

	Context("check error group", func() {
		It("should error", func() {
			defer GinkgoRecover()

			if mirs, e := TagMirFrom(&errGroupEntry{group: "v3"}); e == nil {
				tagMirs = mirs
				err = nil
			} else {
				err = e
			}
			Expect(err).Should(BeNil())
		})
	})

	Context("check error no method", func() {
		It("should error", func() {
			if mirs, e := TagMirFrom(&errNoMethodEntry{}); e == nil {
				tagMirs = mirs
				err = nil
			} else {
				err = e
			}
			Expect(err).ShouldNot(BeNil())
		})
	})
})

func assertTagFields(fields []*TagField) map[string]*TagField {
	pathFields := make(map[string]*TagField, len(fields))
	for _, field := range fields {
		pathFields[field.Path] = field
		handler, ok := field.Handler.(func() string)
		Expect(ok).Should(BeTrue())
		reuslt := handler()
		Expect(reuslt).To(Equal(field.Path))

		switch field.Path {
		case "/query/":
			Expect(field.Queries).Should(HaveLen(1))
		case "/full/{other}/{id:[0-9]+}":
			Expect(field.Host).To(Equal("{subdomain}.domain.com:8013"))
			Expect(field.Queries).Should(HaveLen(3))
		}
	}
	return pathFields
}
