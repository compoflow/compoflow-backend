package argo

import (
	"sync"

	"github.com/Lavender-QAQ/microservice-workflows-backend/pkg/executer/common"
	v1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/beevik/etree"
	v1 "k8s.io/api/core/v1"
)

type PythonscriptNode struct {
	Node
	script string
}

func NewPythonscriptNode(id string, script string) *PythonscriptNode {
	return &PythonscriptNode{
		Node: Node{
			id:     id,
			custom: 2,
			in:     []string{},
			out:    []string{},
		},
		script: script,
	}
}

func (node *PythonscriptNode) GenerateTemplate() v1alpha1.Template {
	template := v1alpha1.Template{
		Name: node.GetId(),
		Script: &v1alpha1.ScriptTemplate{
			Container: v1.Container{},
			Source:    node.script,
		},
	}

	if node.HaveInNode() && node.HaveOutNode() {
		template.Outputs.Artifacts = getTemplateArtifactsByOutcome(node.GetId())
		template.Inputs.Artifacts = getTemplateArtifactsByIncome(node.GetInNode())
	} else if node.HaveOutNode() {
		template.Outputs.Artifacts = getTemplateArtifactsByOutcome(node.GetId())
	} else if node.HaveInNode() {
		template.Inputs.Artifacts = getTemplateArtifactsByIncome(node.GetInNode())
	}
	return template
}

// Add pythonscript node to map
func buildPythonscriptNode(e etree.Element, node_wg *sync.WaitGroup) {
	defer node_wg.Done()
	id := e.SelectAttrValue("id", "none")
	script := e.SelectAttrValue("script", "none")
	var node common.NodeInterface = NewPythonscriptNode(id, script)
	mp_mutex.Lock()
	mp[id] = node
	mp_mutex.Unlock()
}
