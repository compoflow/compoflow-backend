package kubernetes

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var namespace string
var KubeConfig []byte
var clientset *kubernetes.Clientset
var RestConf *rest.Config

func Init(restConf *rest.Config, ns string) error {
	var err error
	clientset, err = InitClient(restConf)
	if err != nil {
		return err
	}
	RestConf = restConf
	namespace = ns
	PodAdapter = clientset.CoreV1().Pods(namespace)
	DeploymentAdapter = clientset.AppsV1().Deployments(namespace)
	ServiceAdapter = clientset.CoreV1().Services(namespace)
	return nil
}

func GetNamespace() string {
	return namespace
}

// Initialize the K8S client
func InitClient(restConf *rest.Config) (*kubernetes.Clientset, error) {
	clientset, err := kubernetes.NewForConfig(restConf)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

// Get RestConf
func GetRestConf() *rest.Config {
	return RestConf
}

// Get ClientSet
func GetClientSet() *kubernetes.Clientset {
	return clientset
}
