package util

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
