package executer

import (
	"github.com/go-logr/logr"
	"github.com/tjcadworkflow/backend/pkg/parser"
)

type WorkflowStarter struct {
	Workflow_id string
	dag         *map[string]*parser.Node
	Logger      logr.Logger
}

// Constructor of the workflow initiator
func NewWorkflowStarter(logger logr.Logger, workflow_id string, dag *map[string]*parser.Node) *WorkflowStarter {
	return &WorkflowStarter{
		Workflow_id: workflow_id,
		dag:         dag,
		Logger:      logger,
	}
}

// Enter the information for the DAG and turn the map into a real workflow
func (w *WorkflowStarter) CreateWorkflow() error {
	_ = w.Logger

	// TODO: Execute workflow

	return nil
}
