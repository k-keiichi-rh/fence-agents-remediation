package v1alpha1

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/medik8s/fence-agents-remediation/pkg/validation"
)

var _ = Describe("FenceAgentsRemediation Validation", func() {

	Context("creating FenceAgentsRemediation", func() {

		When("agent name match format and binary", func() {
			It("should be accepted", func() {
				far := getTestFAR(validAgentName)
				_, err := far.ValidateCreate()
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("agent name was not found ", func() {
			It("should be rejected", func() {
				far := getTestFAR(invalidAgentName)
				_, err := far.ValidateCreate()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("unsupported fence agent: %s", invalidAgentName))
			})
		})

		Context("with OutOfServiceTaint strategy", func() {
			var outOfServiceStrategy *FenceAgentsRemediation

			BeforeEach(func() {
				orgValue := validation.IsOutOfServiceTaintSupported
				DeferCleanup(func() { validation.IsOutOfServiceTaintSupported = orgValue })

				outOfServiceStrategy = getFAR(validAgentName, OutOfServiceTaintRemediationStrategy)
			})
			When("out of service taint is supported", func() {
				BeforeEach(func() {
					validation.IsOutOfServiceTaintSupported = true
				})
				It("should be allowed", func() {
					_, err := outOfServiceStrategy.ValidateCreate()
					Expect(err).To(Succeed())
				})
			})
			When("out of service taint is not supported", func() {
				BeforeEach(func() {
					validation.IsOutOfServiceTaintSupported = false
				})
				It("should be denied", func() {
					_, err := outOfServiceStrategy.ValidateCreate()
					Expect(err).To(MatchError(ContainSubstring(outOfServiceTaintUnsupportedMsg)))
				})
			})
		})
	})

	Context("updating FenceAgentsRemediation", func() {
		var oldFAR *FenceAgentsRemediation
		When("agent name match format and binary", func() {
			BeforeEach(func() {
				oldFAR = getTestFAR(invalidAgentName)
			})
			It("should be accepted", func() {
				far := getTestFAR(validAgentName)
				_, err := far.ValidateUpdate(oldFAR)
				Expect(err).ToNot(HaveOccurred())
			})
		})

		When("agent name was not found ", func() {
			BeforeEach(func() {
				oldFAR = getTestFAR(invalidAgentName)
			})
			It("should be rejected", func() {
				far := getTestFAR(invalidAgentName)
				_, err := far.ValidateUpdate(oldFAR)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("unsupported fence agent: %s", invalidAgentName))
			})
		})
		Context("with OutOfServiceTaint strategy", func() {
			var outOfServiceStrategy *FenceAgentsRemediation
			var resourceDeletionStrategy *FenceAgentsRemediation

			BeforeEach(func() {
				orgValue := validation.IsOutOfServiceTaintSupported
				DeferCleanup(func() { validation.IsOutOfServiceTaintSupported = orgValue })

				outOfServiceStrategy = getFAR(validAgentName, OutOfServiceTaintRemediationStrategy)
				resourceDeletionStrategy = getFAR(validAgentName, ResourceDeletionRemediationStrategy)
			})
			When("out of service taint is supported", func() {
				BeforeEach(func() {
					validation.IsOutOfServiceTaintSupported = true
				})
				It("should be allowed", func() {
					_, err := outOfServiceStrategy.ValidateUpdate(resourceDeletionStrategy)
					Expect(err).To(Succeed())
				})
			})
			When("out of service taint is not supported", func() {
				BeforeEach(func() {
					validation.IsOutOfServiceTaintSupported = false
				})
				It("should be denied", func() {
					_, err := outOfServiceStrategy.ValidateUpdate(resourceDeletionStrategy)
					Expect(err).To(MatchError(ContainSubstring(outOfServiceTaintUnsupportedMsg)))
				})
			})
		})
	})
})

func getTestFAR(agentName string) *FenceAgentsRemediation {
	return getFAR(agentName, ResourceDeletionRemediationStrategy)
}

func getFAR(agentName string, strategy RemediationStrategyType) *FenceAgentsRemediation {
	return &FenceAgentsRemediation{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-" + agentName,
		},
		Spec: FenceAgentsRemediationSpec{
			Agent:               agentName,
			RemediationStrategy: strategy,
		},
	}
}
