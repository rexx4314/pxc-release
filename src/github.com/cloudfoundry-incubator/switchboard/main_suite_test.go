package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func TestSwitchboard(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Switchboard Main Suite")
}

var switchboardBinPath string

var _ = SynchronizedBeforeSuite(
	func() []byte {
		switchboardConfig, err := gexec.Build("github.com/pivotal-cf-experimental/switchboard", "-race")
		Ω(err).ShouldNot(HaveOccurred())
		return []byte(switchboardConfig)
	},
	func(switchboardConfig []byte) {
		switchboardBinPath = string(switchboardConfig)
	},
)

var _ = SynchronizedAfterSuite(func() {
}, func() {
	gexec.CleanupBuildArtifacts()
})
