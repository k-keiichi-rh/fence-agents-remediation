package validation

import (
	"fmt"
	"regexp"
	"strconv"

	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	loggerValidation = ctrl.Log.WithName("validation")
	//IsOutOfServiceTaintSupported will be set to true in case OutOfServiceTaint is supported (k8s 1.26 or higher)
	IsOutOfServiceTaintSupported bool
	leadingDigits                = regexp.MustCompile(`^(\d+)`)
)

const (
	//out of service taint strategy const (supported from 1.26)
	minK8sMajorVersionSupportingOutOfServiceTaint = 1
	minK8sMinorVersionSupportingOutOfServiceTaint = 26
)

// RemoveTaint removes taint from the taint list when it is existed, and returns error if it fails in the process
func InitOutOfServiceTaintSupportedFlag(config *rest.Config) error {
	if cs, err := kubernetes.NewForConfig(config); err != nil || cs == nil {
		if cs == nil {
			err = fmt.Errorf("k8s client set is nil")
		}
		loggerValidation.Error(err, "couldn't retrieve k8s client")
		return err
	} else if k8sVersion, err := cs.Discovery().ServerVersion(); err != nil || k8sVersion == nil {
		if k8sVersion == nil {
			err = fmt.Errorf("k8s server version is nil")
		}
		loggerValidation.Error(err, "couldn't retrieve k8s server version")
		return err
	} else {
		return setOutOfTaintSupportedFlag(k8sVersion)
	}
}

func setOutOfTaintSupportedFlag(version *version.Info) error {
	var majorVer, minorVer int
	var err error
	if majorVer, err = strconv.Atoi(version.Major); err != nil {
		loggerValidation.Error(err, "couldn't parse k8s major version", "major version", version.Major)
		return err
	}
	if minorVer, err = strconv.Atoi(leadingDigits.FindString(version.Minor)); err != nil {
		loggerValidation.Error(err, "couldn't parse k8s minor version", "minor version", version.Minor)
		return err
	}

	IsOutOfServiceTaintSupported = majorVer > minK8sMajorVersionSupportingOutOfServiceTaint || (majorVer == minK8sMajorVersionSupportingOutOfServiceTaint && minorVer >= minK8sMinorVersionSupportingOutOfServiceTaint)
	loggerValidation.Info("out of service taint strategy", "isSupported", IsOutOfServiceTaintSupported, "k8sMajorVersion", majorVer, "k8sMinorVersion", minorVer)
	return nil
}
