package integration_test

import (
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	"testing"
)

var (
	cfg     *rest.Config
	testenv *envtest.Environment
)

func TestMain(m *testing.M) {

}
