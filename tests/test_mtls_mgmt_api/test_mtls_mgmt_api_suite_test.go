// Copyright KhulnaSoft, Ltd.
// Please see the included license file for details.

package test_mtls_mgmt_api

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	ginkgo_util "github.com/khulnasoft/casskube/mage/ginkgo"
	"github.com/khulnasoft/casskube/mage/kubectl"
)

var (
	testName     = "test mtls protecting mgmt api"
	namespace    = "test-mtls-for-mgmt-api"
	dcName       = "dc1"
	dcYaml       = "../testdata/oss-one-node-dc-with-mtls.yaml"
	operatorYaml = "../testdata/operator.yaml"
	dcResource   = fmt.Sprintf("CassandraDatacenter/%s", dcName)
	dcLabel      = fmt.Sprintf("cassandra.khulnasoft.com/datacenter=%s", dcName)
	ns           = ginkgo_util.NewWrapper(testName, namespace)
)

func TestLifecycle(t *testing.T) {
	AfterSuite(func() {
		logPath := fmt.Sprintf("%s/aftersuite", ns.LogDir)
		kubectl.DumpAllLogs(logPath).ExecV()
		fmt.Printf("\n\tPost-run logs dumped at: %s\n\n", logPath)
		ns.Terminate()
	})

	RegisterFailHandler(Fail)
	RunSpecs(t, testName)
}

var _ = Describe(testName, func() {
	Context("when in a new cluster", func() {
		Specify("the operator can start, scale up, and terminate a datacenter where the mgmt api is behind mtls", func() {
			By("creating a namespace")
			err := kubectl.CreateNamespace(namespace).ExecV()
			Expect(err).ToNot(HaveOccurred())

			step := "setting up casskube resources via helm chart"
			ns.HelmInstall("../../charts/casskube-chart")

			ns.WaitForOperatorReady()

			// jam in secrets
			step = "creating mtls secrets"
			k := kubectl.ApplyFiles(
				"../testdata/mtls-certs-server.yaml",
				"../testdata/mtls-certs-client.yaml",
			).InNamespace(namespace)
			ns.ExecAndLog(step, k)

			step = "creating a datacenter resource with 1 rack/1 node"
			k = kubectl.ApplyFiles(dcYaml)
			ns.ExecAndLog(step, k)

			// This takes a while sometimes in my dev environment
			ns.WaitForDatacenterReadyWithTimeouts(dcName, 600, 120)

			step = "scale up to 2 nodes"
			json := "{\"spec\": {\"size\": 2}}"
			k = kubectl.PatchMerge(dcResource, json)
			ns.ExecAndLog(step, k)

			ns.WaitForDatacenterOperatorProgress(dcName, "Updating", 30)
			ns.WaitForDatacenterReady(dcName)

			step = "deleting the dc"
			k = kubectl.DeleteFromFiles(dcYaml)
			ns.ExecAndLog(step, k)

			// TODO FIXME: re-enable this when the following issue is fixed:
			// https://github.com/khulnasoft/management-api-for-apache-cassandra/issues/42
			// k = kubectl.Logs().
			//         WithLabel("statefulset.kubernetes.io/pod-name=cluster1-dc1-r1-sts-0").
			//         WithFlag("container", "cassandra")
			// ns.WaitForOutputContainsAndLog(step, k, "node/drain status=200 OK", 30)

			step = "checking that the dc no longer exists"
			json = "jsonpath={.items}"
			k = kubectl.Get("CassandraDatacenter").
				WithLabel(dcLabel).
				FormatOutput(json)
			ns.WaitForOutputAndLog(step, k, "[]", 300)
		})
	})
})
