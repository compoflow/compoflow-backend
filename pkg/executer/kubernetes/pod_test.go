package kubernetes

import "testing"

func TestGetPodByName(t *testing.T) {
	pod, err := GetPodByName("httpserver-6c7c9c45f7-5tkbg")
	if err != nil {
		t.Error(err)
	}
	t.Log(pod)
}
