package executer

import (
	"github.com/Lavender-QAQ/microservice-workflows-backend/executer/argo"
	"github.com/Lavender-QAQ/microservice-workflows-backend/executer/common"
	"github.com/go-logr/logr"
)

type WorkflowStarter struct {
	Workflow_id string
	dag         *map[string]common.NodeInterface
	Logger      logr.Logger
}

// Constructor of the workflow initiator
func NewWorkflowStarter(logger logr.Logger, workflow_id string, dag *map[string]common.NodeInterface) *WorkflowStarter {
	return &WorkflowStarter{
		Workflow_id: workflow_id,
		dag:         dag,
		Logger:      logger,
	}
}

// Enter the information for the DAG and turn the map into a real workflow
func (w *WorkflowStarter) CreateWorkflow() error {
	err := argo.CreateWorkflow(w.Logger.WithName("argo"), w.Workflow_id, w.dag)
	if err != nil {
		return err
	}
	return nil
}
