package defaulter_handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/erkanzileli/admission-webhooks-the-easy-way/internal/consts"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// PodDefaulterHandler sets the consts.PodAnnotationKey as consts.PodAnnotationValue
type PodDefaulterHandler struct {
	decoder *admission.Decoder
}

func NewPodDefaulterHandler() *PodDefaulterHandler {
	return &PodDefaulterHandler{}
}

func (h *PodDefaulterHandler) Handle(ctx context.Context, req admission.Request) admission.Response {
	pod := &corev1.Pod{}

	err := h.decoder.Decode(req, pod)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	if pod.Annotations == nil {
		pod.Annotations = map[string]string{}
	}
	pod.Annotations[consts.PodAnnotationKey] = consts.PodAnnotationValue

	marshaledPod, err := json.Marshal(pod)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError, err)
	}

	return admission.PatchResponseFromRaw(req.Object.Raw, marshaledPod)
}

// InjectDecoder injects the decoder.
// PodDefaulterHandler implements admission.DecoderInjector so a decoder will be automatically injected.
func (h *PodDefaulterHandler) InjectDecoder(d *admission.Decoder) error {
	h.decoder = d
	return nil
}
