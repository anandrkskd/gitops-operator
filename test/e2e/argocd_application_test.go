package e2e

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/redhat-developer/gitops-operator/test/helper"
)

var _ = Describe("Argo CD metrics controller", func() {
	BeforeEach(func() {
		Expect(helper.CreateNamespace(k8sClient, helper.RandomString(6))).To(BeNil())
	})
	Context("Check if monitoring resources are created", func() {
		It("1-004_validate_argocd_installation", Label("sequential"), func() {

		})

		It("1-009_validate-manage-other-namespace", Label("parallel"), func() {

		})
	})
})
