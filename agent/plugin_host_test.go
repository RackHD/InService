package plugins_test

import (
	. "github.com/RackHD/inservice/agent"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PluginHost", func() {
	It("passes", func() {
		NewPlugin(nil)
		Expect(true).To(BeTrue())
	})
})
