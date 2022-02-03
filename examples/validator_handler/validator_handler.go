package validator_handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/erkanzileli/admission-webhooks-the-easy-way/internal/consts"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// PodValidatorHandler validates consts.PodAnnotationKey annotation
type PodValidatorHandler struct {
	Client  client.Client
	decoder *admission.Decoder
}

func NewPodValidatorHandler() *PodValidatorHandler {
	return &PodValidatorHandler{}
}

// Handle admits a pod if a specific annotation exists.
func (h *PodValidatorHandler) Handle(ctx context.Context, req admission.Request) admission.Response {
	pod := &corev1.Pod{}

	err := h.decoder.Decode(req, pod)
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
// PodValidatorHandler implements admission.DecoderInjector so a decoder will be automatically injected.
func (h *PodValidatorHandler) InjectDecoder(d *admission.Decoder) error {
	h.decoder = d
	return nil
}
