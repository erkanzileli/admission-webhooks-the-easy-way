package defaulter_handler

import (
	"context"
	"encoding/json"
	"github.com/erkanzileli/admission-webhooks-the-easy-way/examples/consts"
	"net/http"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// podMutatingWebhook annotates Pods
type podMutatingWebhook struct {
	decoder *admission.Decoder
}

func NewPodMutatingWebhook() *podMutatingWebhook {
	return &podMutatingWebhook{}
}

// Handle adds an annotation to every incoming pods.
func (a *podMutatingWebhook) Handle(ctx context.Context, req admission.Request) admission.Response {
	pod := &corev1.Pod{}

	err := a.decoder.Decode(req, pod)
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
// podMutatingWebhook implements admission.DecoderInjector so a decoder will be automatically injected.
func (a *podMutatingWebhook) InjectDecoder(d *admission.Decoder) error {
	a.decoder = d
	return nil
}
