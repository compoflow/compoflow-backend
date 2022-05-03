package argo

import (
	"errors"

	"github.com/Lavender-QAQ/microservice-workflows-backend/executer/common"
	v1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/go-logr/logr"
)

func GenerateTemplate(logger logr.Logger, node common.NodeInterface) (v1alpha1.Template, error) {
	switch node.GetCustom() {
	case 1:
		return generateDockerTemplate(node), nil
	case 2:
		return generatePythonscriptTemplate(node), nil
	case 3:
		return generateSuspendTemplate(node), nil
	}
	return v1alpha1.Template{}, errors.New("Node type do not match")
}

func generateDockerTemplate(node common.NodeInterface) v1alpha1.Template {
	return v1alpha1.Template{}
}

func generatePythonscriptTemplate(node common.NodeInterface) v1alpha1.Template {
	return v1alpha1.Template{}
}

func generateSuspendTemplate(node common.NodeInterface) v1alpha1.Template {
	return v1alpha1.Template{}
}
