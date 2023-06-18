package utils

// Copy paste from SNR - https://github.com/medik8s/self-node-remediation/blob/main/pkg/utils/taints.go
import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/medik8s/fence-agents-remediation/api/v1alpha1"
)

var (
	loggerTaint = ctrl.Log.WithName("taints")
)

// taintExists checks if the given taint exists in list of taints. Returns true if exists false otherwise.
func taintExists(taints []corev1.Taint, taintToFind *corev1.Taint) bool {
	for _, taint := range taints {
		if taint.MatchTaint(taintToFind) {
			return true
		}
	}
	return false
}

// deleteTaint removes all the taints that have the same key and effect to given taintToDelete.
func deleteTaint(taints []corev1.Taint, taintToDelete *corev1.Taint) ([]corev1.Taint, bool) {
	var newTaints []corev1.Taint
	deleted := false
	for i := range taints {
		if taintToDelete.MatchTaint(&taints[i]) {
			deleted = true
			continue
		}
		newTaints = append(newTaints, taints[i])
	}
	return newTaints, deleted
}

// createNoExecuteTaint returns a remediation NoExeucte taint
func createNoExecuteTaint() corev1.Taint {
	return corev1.Taint{
		Key:    v1alpha1.Medik8sRemediationTaintKey,
		Value:  v1alpha1.FARRemediationTaintValue,
		Effect: corev1.TaintEffectNoExecute,
	}
}

// AppendTaint appends new taint to the taint list when it is not present, and returns error if it fails in the process
func AppendTaint(r client.Client, nodeName string) error {
	// find node by name
	node, err := getNodeWithName(r, nodeName)
	if err != nil {
		return err
	}

	taint := createNoExecuteTaint()
	// check if taint doesn't exist
	if taintExists(node.Spec.Taints, &taint) {
		return nil
	}
	// add the taint to the taint list
	patch := client.MergeFrom(node.DeepCopy())
	now := metav1.Now()
	taint.TimeAdded = &now
	node.Spec.Taints = append(node.Spec.Taints, taint)

	// update with new taint list
	if err := r.Patch(context.Background(), node, patch); err != nil {
		loggerTaint.Error(err, "Failed to append taint on node", "node name", node.Name, "taint key", taint.Key, "taint effect", taint.Effect)
		return err
	}
	loggerTaint.Info("Taint was added", "taint effect", taint.Effect, "taint list", node.Spec.Taints)
	return nil
}

// RemoveTaint removes taint from the taint list when it is existed, and returns error if it fails in the process
func RemoveTaint(r client.Client, nodeName string) error {
	// find node by name
	node, err := getNodeWithName(r, nodeName)
	if err != nil {
		return err
	}

	taint := createNoExecuteTaint()
	// check if taint exist
	if !taintExists(node.Spec.Taints, &taint) {
		return nil
	}

	// delete the taint from the taint list
	patch := client.MergeFrom(node.DeepCopy())
	if taints, deleted := deleteTaint(node.Spec.Taints, &taint); !deleted {
		loggerTaint.Info("Failed to remove taint from node - taint was not found", "node name", node.Name, "taint key", taint.Key, "taint effect", taint.Effect)
		return nil
	} else {
		node.Spec.Taints = taints
	}

	// update with new taint list
	if err := r.Patch(context.Background(), node, patch); err != nil {
		loggerTaint.Error(err, "Failed to remove taint from node,", "node name", node.Name, "taint key", taint.Key, "taint effect", taint.Effect)
		return err
	}
	loggerTaint.Info("Taint was removed", "taint effect", taint.Effect, "taint list", node.Spec.Taints)
	return nil
}
