/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/medik8s/fence-agents-remediation/pkg/validation"
)

var (
	// webhookTemplateLog is for logging in this package.
	webhookTemplateLog = logf.Log.WithName("fenceagentsremediationtemplate-resource")
)

func (r *FenceAgentsRemediationTemplate) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/validate-fence-agents-remediation-medik8s-io-v1alpha1-fenceagentsremediationtemplate,mutating=false,failurePolicy=fail,sideEffects=None,groups=fence-agents-remediation.medik8s.io,resources=fenceagentsremediationtemplates,verbs=create;update,versions=v1alpha1,name=vfenceagentsremediationtemplate.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &FenceAgentsRemediationTemplate{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *FenceAgentsRemediationTemplate) ValidateCreate() (admission.Warnings, error) {
	webhookTemplateLog.Info("validate create", "name", r.Name)
	return validateStrategy(r.Spec.Template.Spec)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *FenceAgentsRemediationTemplate) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	webhookTemplateLog.Info("validate update", "name", r.Name)
	return validateStrategy(r.Spec.Template.Spec)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *FenceAgentsRemediationTemplate) ValidateDelete() (admission.Warnings, error) {
	// unused for now, add "delete" when needed to verbs in the kubebuilder annotation above
	webhookTemplateLog.Info("validate delete", "name", r.Name)
	return nil, nil
}

func validateStrategy(farSpec FenceAgentsRemediationSpec) (admission.Warnings, error) {
	if farSpec.RemediationStrategy == OutOfServiceTaintRemediationStrategy && !validation.IsOutOfServiceTaintSupported {
		return nil, fmt.Errorf("%s remediation strategy is not supported at kubernetes version lower than 1.26, please use a different remediation strategy", OutOfServiceTaintRemediationStrategy)
	}
	return nil, nil
}
