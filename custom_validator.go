package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

type customPodValidator struct{}

func NewCustomPodValidator() *admission.Webhook {
	return admission.WithCustomValidator(&corev1.Pod{}, &customPodValidator{})
}

func (v *customPodValidator) ValidateCreate(ctx context.Context, obj runtime.Object) error {
	pod := obj.(*corev1.Pod)
	return ensureAnnotationExist(pod)
}

func (v *customPodValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) error {
	pod := newObj.(*corev1.Pod)
	return ensureAnnotationExist(pod)
}

func (v *customPodValidator) ValidateDelete(ctx context.Context, obj runtime.Object) error {
	return nil
}

func ensureAnnotationExist(pod *corev1.Pod) error {
	anno, found := pod.Annotations[podAnnotationKey]
	if !found {
		return fmt.Errorf("missing annotation %s", podAnnotationKey)
	}
	if anno != podAnnotationValue {
		return fmt.Errorf("annotation %s did not have value %q", podAnnotationKey, "foo")
	}
	return nil
}
