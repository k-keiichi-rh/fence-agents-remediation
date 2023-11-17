package v1alpha1

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/medik8s/fence-agents-remediation/pkg/validation"
)

var _ = Describe("FenceAgentsRemediationTemplate Validation", func() {

	Context("a FART", func() {

		var fartValid *FenceAgentsRemediationTemplate
		var outOfServiceStrategy *FenceAgentsRemediationTemplate

		BeforeEach(func() {
			fartValid = &FenceAgentsRemediationTemplate{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test",
				},
				Spec: FenceAgentsRemediationTemplateSpec{
					Template: FenceAgentsRemediationTemplateResource{
						Spec: FenceAgentsRemediationSpec{
							RemediationStrategy: ResourceDeletionRemediationStrategy,
						},
					},
				},
			}
			outOfServiceStrategy = &FenceAgentsRemediationTemplate{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test",
				},
				Spec: FenceAgentsRemediationTemplateSpec{
					Template: FenceAgentsRemediationTemplateResource{
						Spec: FenceAgentsRemediationSpec{
							RemediationStrategy: OutOfServiceTaintRemediationStrategy,
						},
					},
				},
			}

		})

		Context("with valid strategy", func() {
			It("should be allowed", func() {
				_, err := fartValid.ValidateCreate()
				Expect(err).To(Succeed())
			})
		})

		Context("with out Of Service Taint strategy", func() {
			BeforeEach(func() {
				orgValue := validation.IsOutOfServiceTaintSupported
				DeferCleanup(func() { validation.IsOutOfServiceTaintSupported = orgValue })

			})
			When("out of service taint is supported", func() {
				BeforeEach(func() {
					validation.IsOutOfServiceTaintSupported = true
				})
				It("should be allowed", func() {
					_, err := outOfServiceStrategy.ValidateCreate()
					Expect(err).To(Succeed())
					_, err = fartValid.ValidateUpdate(outOfServiceStrategy)
					Expect(err).To(Succeed())
				})
			})
			When("out of service taint is not supported", func() {
				BeforeEach(func() {
					validation.IsOutOfServiceTaintSupported = false
				})
				It("should be denied", func() {
					_, err := outOfServiceStrategy.ValidateCreate()
					Expect(err).To(MatchError(ContainSubstring("OutOfServiceTaint remediation strategy is not supported at kubernetes version lower than 1.26, please use a different remediation strategy")))
					_, err = outOfServiceStrategy.ValidateUpdate(fartValid)
					Expect(err).To(MatchError(ContainSubstring("OutOfServiceTaint remediation strategy is not supported at kubernetes version lower than 1.26, please use a different remediation strategy")))
				})
			})

		})

	})

})
