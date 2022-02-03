package custom_defaulter_test

import (
	customDefaulter "github.com/erkanzileli/admission-webhooks-the-easy-way/examples/custom_defaulter"
	"github.com/erkanzileli/admission-webhooks-the-easy-way/internal/consts"
	"github.com/erkanzileli/admission-webhooks-the-easy-way/internal/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCustomPodDefaulter(t *testing.T) {
	defaulter := customDefaulter.NewCustomPodDefaulterWebhook()
	assert.NotNil(t, defaulter)
}

func Test_customPodDefaulter_Default_when_annotation_is_not_exist(t *testing.T) {
	// Given
	pod := util.CreatePod()
	defaulter := customDefaulter.CustomPodDefaulter{}

	// When
	err := defaulter.Default(nil, pod)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, consts.PodAnnotationValue, pod.Annotations[consts.PodAnnotationKey])
}

func Test_customPodDefaulter_Default_when_annotation_is_exist_with_different_value(t *testing.T) {
	// Given
	pod := util.CreatePod(consts.PodAnnotationKey, "bar")
	defaulter := customDefaulter.CustomPodDefaulter{}

	// When
	err := defaulter.Default(nil, pod)

	// Then
	assert.NoError(t, err)
	assert.Equal(t, consts.PodAnnotationValue, pod.Annotations[consts.PodAnnotationKey])
}
