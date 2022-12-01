package kubernetes

import (
	"context"

	// appsv1 "k8s.io/api/apps/v1"
	// apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	v1 "k8s.io/api/core/v1"
	type_corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

var PodAdapter type_corev1.PodInterface

func getPodByName(podName string) (*v1.Pod, error) {
	pod, err := PodAdapter.Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return pod, nil
}
