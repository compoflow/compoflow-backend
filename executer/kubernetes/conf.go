package kubernetes

import (
	"io/ioutil"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const namespace string = "argo"

var KubeConfig []byte
var RestConf *rest.Config
var clientset *kubernetes.Clientset

func Init(kubeconfigPath string) error {
	var err error
	clientset, err = InitClient(kubeconfigPath)
	if err != nil {
		return err
	}
	PodAdapter = clientset.CoreV1().Pods(namespace)
	DeploymentAdapter = clientset.AppsV1().Deployments(namespace)
	ServiceAdapter = clientset.CoreV1().Services(namespace)
	return nil
}

func GetNamespace() string {
	return namespace
}

// Initialize the K8S client
func InitClient(filepath string) (*kubernetes.Clientset, error) {
	restConf, err := GetRestConf(filepath)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(restConf)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

// Generate RestConf
func GetRestConf(filepath string) (*rest.Config, error) {
	var err error
	if KubeConfig, err = ioutil.ReadFile(filepath); err != nil {
		return nil, err
	}
	// Generate restful client
	if RestConf, err = clientcmd.RESTConfigFromKubeConfig(KubeConfig); err != nil {
		return nil, err
	}
	return RestConf, nil
}
