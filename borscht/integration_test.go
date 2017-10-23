package borscht_test

import (
	"github.com/craigfurman/borscht/borscht"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("borscht as a library", func() {
	var (
		releasePath string
		fromVersion string
		toVersion   string
	)

	BeforeEach(func() {
		releasePath = envMustHave("BORSCHT_GRR_PATH")
	})

	Describe("success", func() {
		var jobDiffs map[string]string

		JustBeforeEach(func() {
			var err error
			jobDiffs, err = borscht.Diff(releasePath, fromVersion, toVersion)
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

			Context("and the job specs have changed", func() {
				BeforeEach(func() {
					fromVersion = "1.8.0"
					toVersion = "1.9.5"
				})

				It("returns jobs with empty diffs", func() {
					Expect(jobDiffs).To(Equal(map[string]string{
						"garden":         example180to195gardenDiff,
						"garden-windows": example180to196gardenWindowsDiff,
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

const example180to195gardenDiff = `diff --git a/jobs/garden/spec b/jobs/garden/spec
index 5861794..ca5e870 100644
--- a/jobs/garden/spec
+++ b/jobs/garden/spec
@@ -5,7 +5,6 @@ templates:
   garden_ctl.erb: bin/garden_ctl
   garden-default.erb: config/garden-default
   auplink: bin/auplink
-  brutefs: bin/brutefs
 
 packages:
   - guardian
@@ -15,6 +14,12 @@ packages:
   - tar
   - garden-idmapper
 
+provides:
+  - name: rootless_link
+    type: garden_rootless_link
+    properties:
+      - garden.experimental_rootless_mode
+
 properties:
   garden.listen_network:
     description: "Garden server connection mode (tcp or unix)."
`

const example180to196gardenWindowsDiff = `diff --git a/jobs/garden-windows/spec b/jobs/garden-windows/spec
index 5573f81..43a4952 100644
--- a/jobs/garden-windows/spec
+++ b/jobs/garden-windows/spec
@@ -12,5 +12,45 @@ properties:
     description: "Garden server listening address."
     default: 127.0.0.1:7777
 
+  garden.dropsonde.origin:
+    description: "A string identifier that will be used when reporting metrics to Dropsonde."
+    default: "garden-windows"
+
+  garden.dropsonde.destination:
+    description: "A URL that points at the Metron agent to which metrics are forwarded. By default, it matches with the default of Metron."
+
+  garden.log_level:
+    description: "log level for the Garden server - can be debug, info, error or fatal"
+    default: info
+
   garden.runtime_plugin:
     description: "Path to a runtime plugin binary"
+
+  garden.image_plugin:
+    description: "Path to an image plugin binary"
+
+  garden.network_plugin:
+    description: "Path to a network plugin binary"
+
+  garden.network_plugin_extra_args:
+    description: "An array of additional arguments which will be passed to the network plugin binary"
+    default: []
+
+  garden.nstar_bin:
+    description: "Path to nstar binary"
+
+  garden.tar_bin:
+    description: "Path to tar binary"
+    default: "C:\\var\\vcap\\bosh\\bin\\tar.exe"
+
+  garden.max_containers:
+    description: "Maximum container capacity to advertise. It is not recommended to set this larger than 75."
+    default: 75
+
+  garden.destroy_containers_on_start:
+    description: "If true, all existing containers will be destroyed any time the garden server starts up"
+    default: false
+
+  garden.default_container_rootfs:
+    description: "path to the rootfs to use when a container specifies no rootfs"
+    default: ""
`
