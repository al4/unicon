package gonfig_test

import (
	. "github.com/Nomon/gonfig"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ArgvConfig", func() {
	var (
		err error
		cfg Configurable
	)
	BeforeEach(func() {
		cfg = NewArgvConfig("test")
		err = cfg.Load()
	})
	It("Should load variables from commandline", func() {
		Expect(len(cfg.All()) > 0).To(BeTrue())
	})
})