package validator_handler

import (
	"context"
	"fmt"
	"github.com/erkanzileli/admission-webhooks-the-easy-way/internal/consts"
	"net/http"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// CustomPodValidator validates consts.PodAnnotationKey annotation
type podValidatingWebhook struct {
	Client  client.Client
	decoder *admission.Decoder
}

func NewPodValidatingWebhook() *podValidatingWebhook {
	return &podValidatingWebhook{}
}

// Handle admits a pod if a specific annotation exists.
func (v *podValidatingWebhook) Handle(ctx context.Context, req admission.Request) admission.Response {
	pod := &corev1.Pod{}

	err := v.decoder.Decode(req, pod)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	anno, found := pod.Annotations[consts.PodAnnotationKey]
	if !found {
		return admission.Denied(fmt.Sprintf("missing annotation %s", consts.PodAnnotationKey))
	}
	if anno != consts.PodAnnotationValue {
		return admission.Denied(fmt.Sprintf("annotation %s did not have value %q", consts.PodAnnotationKey, "foo"))
	}

	return admission.Allowed("")
}

// InjectDecoder injects the decoder.
// podValidatingWebhook implements admission.DecoderInjector so a decoder will be automatically injected.
func (v *podValidatingWebhook) InjectDecoder(d *admission.Decoder) error {
	v.decoder = d
	return nil
}
