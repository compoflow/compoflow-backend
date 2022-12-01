package kubernetes

import (
	"context"
	"errors"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	// corev1 "k8s.io/client-go/applyconfigurations/core/v1"

	type_appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
)

var DeploymentAdapter type_appsv1.DeploymentInterface

func getDeploymenyByName(deploymentName string) (*appsv1.Deployment, error) {
	deployment, err := DeploymentAdapter.Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return deployment, nil
}

// Create the adaptive Service while creating the Deployment (call CreateService)
func createDeployment(name string, image string, port int, command []string, args []string) error {
	replicas := int32(1)
	labels := make(map[string]string, 1)
	labels["app"] = name
	containers := newContainerQueue()
	containers.push(name, port, image, command, args)
	deployment := appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{MatchLabels: labels},
			Template: v1.PodTemplateSpec{ObjectMeta: metav1.ObjectMeta{Labels: labels}, Spec: v1.PodSpec{Containers: containers}},
		},
		Status: appsv1.DeploymentStatus{},
	}
	_, err := DeploymentAdapter.Create(context.TODO(), &deployment, metav1.CreateOptions{})
	if err != nil {
		return errors.New("创建Deployment错误:" + err.Error())
	}
	err = createService(name, labels, port)
	if err != nil {
		return errors.New("创建Service错误" + err.Error())
	}
	return nil
}

func deleteDeployment(name string) error {
	err := DeploymentAdapter.Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	err = deleteService(name + "-svc")
	if err != nil {
		return err
	}
	return nil
}
