package v1alpha1

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/medik8s/fence-agents-remediation/pkg/validation"
)

var _ = Describe("FenceAgentsRemediationTemplate Validation", func() {

	Context("creating FenceAgentsRemediationTemplate", func() {

		Context("with ResourceDeletion strategy", func() {
			When("agent name match format and binary", func() {
				It("should be accepted", func() {
					farTemplate := getTestFARTemplate(validAgentName)
					Expect(farTemplate.ValidateCreate()).Error().NotTo(HaveOccurred())
				})
			})

			When("agent name was not found ", func() {
				It("should be rejected", func() {
					farTemplate := getTestFARTemplate(invalidAgentName)
					Expect(farTemplate.ValidateCreate()).Error().To(MatchError(ContainSubstring("unsupported fence agent: %s", invalidAgentName)))
				})
			})
		})

		Context("with OutOfServiceTaint strategy", func() {
			var outOfServiceStrategy *FenceAgentsRemediationTemplate

			BeforeEach(func() {
				orgValue := validation.IsOutOfServiceTaintSupported
				DeferCleanup(func() { validation.IsOutOfServiceTaintSupported = orgValue })

				outOfServiceStrategy = getFARTemplate(validAgentName, OutOfServiceTaintRemediationStrategy)
			})

			When("out of service taint is supported", func() {
				BeforeEach(func() {
					validation.IsOutOfServiceTaintSupported = true
				})
				It("should be allowed", func() {
					Expect(outOfServiceStrategy.ValidateCreate()).Error().NotTo(HaveOccurred())
				})
			})

			When("out of service taint is not supported", func() {
				BeforeEach(func() {
					validation.IsOutOfServiceTaintSupported = false
				})
				It("should be denied", func() {
					Expect(outOfServiceStrategy.ValidateCreate()).Error().To(MatchError(ContainSubstring(outOfServiceTaintUnsupportedMsg)))
				})
			})
		})
	})

	Context("updating FenceAgentsRemediationTemplate", func() {

		Context("with ResourceDeletion strategy", func() {
			var oldFARTemplate *FenceAgentsRemediationTemplate

			When("agent name match format and binary", func() {
				BeforeEach(func() {
					oldFARTemplate = getTestFARTemplate(invalidAgentName)
				})
				It("should be accepted", func() {
					farTemplate := getTestFARTemplate(validAgentName)
					Expect(farTemplate.ValidateUpdate(oldFARTemplate)).Error().NotTo(HaveOccurred())
				})
			})

			When("agent name was not found ", func() {
				BeforeEach(func() {
					oldFARTemplate = getTestFARTemplate(invalidAgentName)
				})
				It("should be rejected", func() {
					farTemplate := getTestFARTemplate(invalidAgentName)
					Expect(farTemplate.ValidateUpdate(oldFARTemplate)).Error().To(MatchError(ContainSubstring("unsupported fence agent: %s", invalidAgentName)))
				})
			})
		})

		Context("with OutOfServiceTaint strategy", func() {
			var outOfServiceStrategy *FenceAgentsRemediationTemplate
			var resourceDeletionStrategy *FenceAgentsRemediationTemplate

			BeforeEach(func() {
				orgValue := validation.IsOutOfServiceTaintSupported
				DeferCleanup(func() { validation.IsOutOfServiceTaintSupported = orgValue })

				outOfServiceStrategy = getFARTemplate(validAgentName, OutOfServiceTaintRemediationStrategy)
				resourceDeletionStrategy = getFARTemplate(validAgentName, ResourceDeletionRemediationStrategy)
			})

			When("out of service taint is supported", func() {
				BeforeEach(func() {
					validation.IsOutOfServiceTaintSupported = true
				})
				It("should be allowed", func() {
					Expect(outOfServiceStrategy.ValidateUpdate(resourceDeletionStrategy)).Error().NotTo(HaveOccurred())
				})
			})

			When("out of service taint is not supported", func() {
				BeforeEach(func() {
					validation.IsOutOfServiceTaintSupported = false
				})
				It("should be denied", func() {
					Expect(outOfServiceStrategy.ValidateUpdate(resourceDeletionStrategy)).Error().To(MatchError(ContainSubstring(outOfServiceTaintUnsupportedMsg)))
				})
			})
		})
	})
})

func getTestFARTemplate(agentName string) *FenceAgentsRemediationTemplate {
	return getFARTemplate(agentName, ResourceDeletionRemediationStrategy)
}

func getFARTemplate(agentName string, strategy RemediationStrategyType) *FenceAgentsRemediationTemplate {
	return &FenceAgentsRemediationTemplate{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-" + agentName + "-template",
		},
		Spec: FenceAgentsRemediationTemplateSpec{
			Template: FenceAgentsRemediationTemplateResource{
				Spec: FenceAgentsRemediationSpec{
					Agent:               agentName,
					RemediationStrategy: strategy,
				},
			},
		},
	}
}
