package kubernetes

import "testing"

func TestGetServiceByName(t *testing.T) {
	svc, err := GetServiceByName("httpserver-svc")
	if err != nil {
		t.Error(err)
	}
	t.Log(svc)
}
