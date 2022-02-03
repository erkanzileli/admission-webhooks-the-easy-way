package validator_handler_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/erkanzileli/admission-webhooks-the-easy-way/examples/validator_handler"
	"github.com/erkanzileli/admission-webhooks-the-easy-way/internal/consts"
	"github.com/erkanzileli/admission-webhooks-the-easy-way/internal/util"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var (
	scheme  = runtime.NewScheme()
	decoder *admission.Decoder
)

func TestMain(m *testing.M) {
	dcd, err := admission.NewDecoder(scheme)
	if err != nil {
		log.Printf("error while creating decoder")
		panic(err)
	}
	decoder = dcd

	m.Run()
}

func TestNewPodValidatorHandler(t *testing.T) {
	handler := validator_handler.NewPodValidatorHandler()
	assert.NotNil(t, handler)
}

func TestPodValidatorHandler_Handle_when_annotation_is_exist(t *testing.T) {
	// Given
	handler := validator_handler.NewPodValidatorHandler()
	handler.InjectDecoder(decoder)
	pod := util.CreatePod(consts.PodAnnotationKey, consts.PodAnnotationValue)
	req := newAdmissionReq(pod)

	// When
	response := handler.Handle(nil, req)

	// Then
	assert.True(t, response.Allowed)
}

func TestPodValidatorHandler_Handle_when_annotation_is_not_exist(t *testing.T) {
	// Given
	handler := validator_handler.NewPodValidatorHandler()
	handler.InjectDecoder(decoder)
	pod := util.CreatePod()
	req := newAdmissionReq(pod)

	// When
	response := handler.Handle(nil, req)

	// Then
	assert.False(t, response.Allowed)
}

func TestPodValidatorHandler_Handle_when_annotation_is_different(t *testing.T) {
	// Given
	handler := validator_handler.NewPodValidatorHandler()
	handler.InjectDecoder(decoder)
	pod := util.CreatePod(consts.PodAnnotationKey, "bar")
	req := newAdmissionReq(pod)

	// When
	response := handler.Handle(nil, req)

	// Then
	assert.False(t, response.Allowed)
}

func TestPodValidatorHandler_Handle_when_pod_is_nil(t *testing.T) {
	// Given
	handler := validator_handler.NewPodValidatorHandler()
	handler.InjectDecoder(decoder)
	req := newAdmissionReq(nil)
	req.Object.Raw = []byte("*")

	// When
	response := handler.Handle(nil, req)

	// Then
	assert.False(t, response.Allowed)
	assert.EqualValues(t, http.StatusBadRequest, response.Result.Code)
}

func newAdmissionReq(pod *corev1.Pod) admission.Request {
	rawPod, _ := json.Marshal(pod)
	return admission.Request{
		AdmissionRequest: admissionv1.AdmissionRequest{
			Object: runtime.RawExtension{
				Raw: rawPod,
			},
		},
	}
}
