package router

import (
	"net/http"

	"github.com/go-logr/logr"
	"github.com/tjcadworkflow/backend/pkg/handlers"
)

var RouterLogger logr.Logger

func NewRouter(ip_port string) {
	logger := RouterLogger

	http.HandleFunc("/deploy", handlers.DeployHandler)
	http.HandleFunc("/save", handlers.SaveHandler)

	logger.WithValues("listening ip address and port", ip_port).Info("Start listen")

	err := http.ListenAndServe(ip_port, nil)
	if err != nil {
		logger.Error(err, "Router listen error")
		return
	}
}
