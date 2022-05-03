package kubernetes

import "testing"

func TestInitClient(t *testing.T) {
	err := Init("../../kubeconfig")
	if err != nil {
		t.Error(err)
		return
	}
}
