package borscht_test

import (
	"github.com/craigfurman/borscht/borscht"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("borscht as a library", func() {
	var (
		b           *borscht.Borscht
		releasePath string
		fromVersion string
		toVersion   string
	)

	BeforeEach(func() {
		b = &borscht.Borscht{}
		releasePath = envMustHave("BORSCHT_GRR_PATH")
	})

	Describe("success", func() {
		var jobDiffs map[string]string

		JustBeforeEach(func() {
			var err error
			jobDiffs, err = b.Diff(releasePath, fromVersion, toVersion)
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when the release path looks like a bosh release", func() {
			Context("and the versions are identical", func() {
				BeforeEach(func() {
					fromVersion = "1.9.5"
					toVersion = fromVersion
				})

				It("returns jobs with empty diffs", func() {
					Expect(jobDiffs).To(Equal(map[string]string{
						"garden":         "",
						"garden-windows": "",
					}))
				})
			})
		})
	})

	Describe("failure", func() {
		PContext("when one of the release versions does not exist", func() {})

		PContext("when the release path doesn't contain a jobs dir", func() {})

		PContext("when the release path doesn't exist", func() {})
	})
})
