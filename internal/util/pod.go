package util

import (
	"encoding/json"
	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

func CreatePod(annotations ...string) (pod *corev1.Pod) {
	pod = &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "test-pod",
			Annotations: map[string]string{},
		},
	}

	if len(annotations) < 2 {
		return
	}

	var annKey string
	for i, anno := range annotations {
		if i%2 == 0 {
			annKey = anno
		}
		pod.Annotations[annKey] = anno
	}

	return
}

func NewAdmissionReq(pod *corev1.Pod) admission.Request {
	rawPod, _ := json.Marshal(pod)
	return admission.Request{
		AdmissionRequest: admissionv1.AdmissionRequest{
			Object: runtime.RawExtension{
				Raw: rawPod,
			},
		},
	}
}
