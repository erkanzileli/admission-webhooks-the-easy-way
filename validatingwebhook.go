package main

import (
	"context"
	"fmt"
	"net/http"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// podValidatingWebhook validates Pods
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

	anno, found := pod.Annotations[podAnnotationKey]
	if !found {
		return admission.Denied(fmt.Sprintf("missing annotation %s", podAnnotationKey))
	}
	if anno != podAnnotationValue {
		return admission.Denied(fmt.Sprintf("annotation %s did not have value %q", podAnnotationKey, "foo"))
	}

	return admission.Allowed("")
}

// podValidatingWebhook implements admission.DecoderInjector.
// A decoder will be automatically injected.

// InjectDecoder injects the decoder.
func (v *podValidatingWebhook) InjectDecoder(d *admission.Decoder) error {
	v.decoder = d
	return nil
}
