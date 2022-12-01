package kubernetes

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	type_v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

var ServiceAdapter type_v1.ServiceInterface

func getServiceByName(serviceName string) (*v1.Service, error) {
	service, err := ServiceAdapter.Get(context.TODO(), serviceName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return service, nil
}

func createService(name string, labels map[string]string, port int) error {
	var servicePorts []v1.ServicePort
	servicePort := v1.ServicePort{
		Protocol:   v1.ProtocolTCP,
		Port:       int32(port),
		TargetPort: intstr.FromInt(port),
	}
	servicePorts = append(servicePorts, servicePort)
	service := v1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name + "-svc",
			Namespace: namespace,
		},
		Spec: v1.ServiceSpec{
			Type:     "ClusterIP",
			Selector: labels,
			Ports:    servicePorts,
		},
	}
	_, err := ServiceAdapter.Create(context.TODO(), &service, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func deleteService(name string) error {
	err := ServiceAdapter.Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}
