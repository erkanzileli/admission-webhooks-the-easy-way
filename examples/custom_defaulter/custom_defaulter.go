package custom_defaulter

import (
	"context"
	"github.com/erkanzileli/admission-webhooks-the-easy-way/internal/consts"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// CustomPodDefaulter sets the consts.PodAnnotationKey as consts.PodAnnotationValue
type CustomPodDefaulter struct{}

func NewCustomPodDefaulterWebhook() *admission.Webhook {
	return admission.WithCustomDefaulter(&corev1.Pod{}, &CustomPodDefaulter{})
}

func (CustomPodDefaulter) Default(ctx context.Context, obj runtime.Object) error {
	pod := obj.(*corev1.Pod)

	logrus.WithField("name", pod.Name).Infof("got defaulting request for pod")

	if pod.Annotations == nil {
		pod.Annotations = map[string]string{}
	}

	pod.Annotations[consts.PodAnnotationKey] = consts.PodAnnotationValue

	return nil
}
