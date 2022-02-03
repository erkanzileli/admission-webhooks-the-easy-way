package custom_validator

import (
	"context"
	"fmt"

	"github.com/erkanzileli/admission-webhooks-the-easy-way/internal/consts"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// CustomPodValidator validates consts.PodAnnotationKey annotation
type CustomPodValidator struct{}

func NewCustomPodValidatorWebhook() *admission.Webhook {
	return admission.WithCustomValidator(&corev1.Pod{}, &CustomPodValidator{})
}

func (CustomPodValidator) ValidateCreate(ctx context.Context, obj runtime.Object) error {
	pod := obj.(*corev1.Pod)
	anno, found := pod.Annotations[consts.PodAnnotationKey]
	if !found {
		return fmt.Errorf("missing annotation %s", consts.PodAnnotationKey)
	}
	if anno != consts.PodAnnotationValue {
		return fmt.Errorf("annotation %s did not have value %q", consts.PodAnnotationKey, consts.PodAnnotationValue)
	}
	return nil
}

func (CustomPodValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) error {
	oldPod := oldObj.(*corev1.Pod)
	newPod := newObj.(*corev1.Pod)

	if oldPod.Annotations[consts.PodAnnotationKey] != newPod.Annotations[consts.PodAnnotationKey] {
		return fmt.Errorf("you can't change annotation %s once you created it", consts.PodAnnotationKey)
	}

	return nil
}

func (CustomPodValidator) ValidateDelete(ctx context.Context, obj runtime.Object) error {
	return nil
}
