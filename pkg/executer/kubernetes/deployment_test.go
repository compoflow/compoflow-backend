package kubernetes

import "testing"

func TestGetDeploymentByName(t *testing.T) {
	dep, err := GetDeploymenyByName("httpserver")
	if err != nil {
		t.Error(err)
	}
	t.Log(dep)
}

func TestCreateDeployment(t *testing.T) {
	command := []string{"/bin/sh"}
	args := []string{"-c", "./httpserver"}
	err := CreateDeployment("httpserver-auto", "192.168.31.60/argo/httpserver:test",
		8080, command, args)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestDeleteDeployment(t *testing.T) {
	err := DeleteDeployment("httpserver-auto")
	if err != nil {
		t.Error(err)
		return
	}
}
