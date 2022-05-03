package argo

import (
	"sync"

	"github.com/Lavender-QAQ/microservice-workflows-backend/executer/common"
	v1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/beevik/etree"
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

func (p *PythonscriptNode) GenerateTemplate() v1alpha1.Template {
	return v1alpha1.Template{}
}
