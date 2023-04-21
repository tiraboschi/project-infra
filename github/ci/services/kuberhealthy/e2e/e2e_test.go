package e2e

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"kubevirt.io/project-infra/github/ci/services/common/k8s/pkg/check"
)

const (
	testNamespace = "kuberhealthy"
)

func TestKuberhealthyDeployment(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "kuberhealthy deployment suite")
}

var _ = Describe("kuberhealthy deployment", func() {
	DescribeTable("should deploy HTTP services",
		func(svcPort, labelSelector, expectedContent, urlPath string) {
			check.HTTPService(testNamespace, svcPort, labelSelector, expectedContent, urlPath)
		},
		Entry("kuberhealthy", "8080", "app=kuberhealthy", `"OK": true`, ""),
	)
})
