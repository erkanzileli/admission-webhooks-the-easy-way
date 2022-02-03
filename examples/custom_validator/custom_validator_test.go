package custom_validator_test

import (
	customValidator "github.com/erkanzileli/admission-webhooks-the-easy-way/examples/custom_validator"
	"github.com/erkanzileli/admission-webhooks-the-easy-way/internal/consts"
	"github.com/erkanzileli/admission-webhooks-the-easy-way/internal/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCustomPodValidator(t *testing.T) {
	validator := customValidator.NewCustomPodValidatorWebhook()
	assert.NotNil(t, validator)
}

func Test_CustomPodValidator_ValidateCreate_when_annotation_is_exist(t *testing.T) {
	// Given
	pod := util.CreatePod(consts.PodAnnotationKey, consts.PodAnnotationValue)
	validator := customValidator.CustomPodValidator{}

	// When
	err := validator.ValidateCreate(nil, pod)

	// Then
	assert.NoError(t, err)
}

func Test_CustomPodValidator_ValidateCreate_when_annotation_is_not_exist(t *testing.T) {
	// Given
	pod := util.CreatePod()
	validator := customValidator.CustomPodValidator{}

	// When
	err := validator.ValidateCreate(nil, pod)

	// Then
	assert.Error(t, err)
}

func Test_CustomPodValidator_ValidateCreate_when_annotation_is_as_expected(t *testing.T) {
	// Given
	pod := util.CreatePod(consts.PodAnnotationKey, "bar")
	validator := customValidator.CustomPodValidator{}

	// When
	err := validator.ValidateCreate(nil, pod)

	// Then
	assert.Error(t, err)
}

func Test_CustomPodValidator_ValidateUpdate_when_annotations_is_same(t *testing.T) {
	// Given
	oldPod := util.CreatePod(consts.PodAnnotationKey, consts.PodAnnotationValue)
	newPod := util.CreatePod(consts.PodAnnotationKey, consts.PodAnnotationValue)
	validator := customValidator.CustomPodValidator{}

	// When
	err := validator.ValidateUpdate(nil, oldPod, newPod)

	// Then
	assert.NoError(t, err)
}

func Test_CustomPodValidator_ValidateUpdate_when_annotation_is_changed(t *testing.T) {
	// Given
	oldPod := util.CreatePod(consts.PodAnnotationKey, consts.PodAnnotationValue)
	newPod := util.CreatePod(consts.PodAnnotationKey, "bar")
	validator := customValidator.CustomPodValidator{}

	// When
	err := validator.ValidateUpdate(nil, oldPod, newPod)

	// Then
	assert.Error(t, err)
}
