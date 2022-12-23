package main

import (
	"flag"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/compoflow/compoflow-backend/pkg/handlers"
	"github.com/compoflow/compoflow-backend/pkg/kubernetes"
	"github.com/compoflow/compoflow-backend/pkg/router"
)

var (
	logger    = ctrl.Log.WithName("workflow")
	namespace string
	listen    string
)

func registerLogger() error {

	// Register handler
	handlers.HandlerLogger = logger.WithName("Handler")
	router.RouterLogger = logger.WithName("Router")

	return nil
}

func main() {
	// Customize flags
	flag.StringVar(&namespace, "namespace", "default", "Specify a namespace for the workflow to run")
	flag.StringVar(&listen, "listen", "0.0.0.0:8080", "Specify the listening ip address and port")

	opts := zap.Options{
		Development: true,
	}
	opts.BindFlags(flag.CommandLine)
	flag.Parse()

	// Get restconfig
	restConf := ctrl.GetConfigOrDie()

	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	// Register logger
	err := registerLogger()
	if err != nil {
		logger.Error(err, "Fail to register logger")
		return
	}

	// Init client-go
	err = kubernetes.Init(restConf, namespace)
	if err != nil {
		logger.Error(err, "Fail to initialize kubernetes cluster")
		return
	}

	// Launch Router
	router.NewRouter(listen)
}
