/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"os"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/joho/godotenv"

	"github.com/Lavender-QAQ/microservice-workflows-backend/pkg/executer/kubernetes"
	"github.com/Lavender-QAQ/microservice-workflows-backend/pkg/handler"
	"github.com/Lavender-QAQ/microservice-workflows-backend/pkg/router"
)

var (
	logger    = ctrl.Log.WithName("run")
	namespace string
	listen    string
)

func init() {
	// 从本地读取环境变量
	_ = godotenv.Load()
	if os.Getenv("ACTIVE_ENV") == "DEV" {
		_ = godotenv.Load(".env.dev")
	} else if os.Getenv("ACTIVE_ENV") == "PROD" {
		_ = godotenv.Load(".env.prod")
	}
}

func registerLogger() error {

	// Register handler
	handler.HandlerLogger = logger.WithName("Handler")
	router.RouterLogger = logger.WithName("Router")

	return nil
}

func main() {
	// Customize flags
	flag.StringVar(&namespace, "namespace", "argo", "Specify a namespace for the workflow to run")
	flag.StringVar(&listen, "listen", "127.0.0.1:30086", "Specify the listening ip address and port")

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
