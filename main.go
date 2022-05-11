package main

import (
	"flag"
	"fmt"
	"github.com/Lavender-QAQ/microservice-workflows-backend/conf"

	"github.com/Lavender-QAQ/microservice-workflows-backend/executer/kubernetes"
	"github.com/Lavender-QAQ/microservice-workflows-backend/handler"
	"github.com/Lavender-QAQ/microservice-workflows-backend/router"
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

var logger logr.Logger

func main() {
	conf.Init()

	kubeconfigPath := flag.String("kubeconfig", "./kubeconfig", "Kubernetes configuration file location")
	listen := flag.String("listen", "127.0.0.1:30086", "Specify the listening ip address and port")
	flag.Parse()

	err := registerLogger()
	if err != nil {
		fmt.Println(err)
		return
	}

	logger.WithValues("kubeconfig location", *kubeconfigPath).Info("Kubeconfig parameters were successfully parsed")

	err = kubernetes.Init(*kubeconfigPath)
	if err != nil {
		logger.Error(err, "Fail to initialize kubernetes cluster")
		return
	}

	err = router.NewRouter(*listen)
	if err != nil {
		logger.Error(err, "Fail to create router")
		return
	}
}

func registerLogger() error {
	zapLog, err := zap.NewDevelopment()
	if err != nil {
		return fmt.Errorf("who watches the watchmen (%v)?", err)
	}
	logger = zapr.NewLogger(zapLog)

	// Register handler
	handler.HandlerLogger = logger.WithName("Handler")
	router.RouterLogger = logger.WithName("Router")

	return nil
}
