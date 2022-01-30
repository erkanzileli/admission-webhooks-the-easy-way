package main

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type customPodDefaulter struct{}

func NewCustomPodDefaulter() *admission.Webhook {
	return admission.WithCustomDefaulter(&corev1.Pod{}, &customPodDefaulter{})
}

func (d *customPodDefaulter) Default(ctx context.Context, obj runtime.Object) error {
	pod := obj.(*corev1.Pod)
	if pod.Annotations == nil {
		pod.Annotations = map[string]string{}
	}
	pod.Annotations[podAnnotationKey] = podAnnotationValue
	return nil
}
