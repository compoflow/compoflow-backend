package argo

import (
	"strconv"
	"sync"

	"github.com/Lavender-QAQ/microservice-workflows-backend/pkg/executer/common"
	v1alpha1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	"github.com/beevik/etree"
)

type SuspendNode struct {
	Node
	suspend int
}

func NewSuspendNode(id string, suspend int) *SuspendNode {
	return &SuspendNode{
		Node: Node{
			id:     id,
			custom: 3,
			in:     []string{},
			out:    []string{},
		},
		suspend: suspend,
	}
}

func (s *SuspendNode) GenerateTemplate() v1alpha1.Template {
	return v1alpha1.Template{}
}

// Add suspend node to map
func buildSuspendNode(e etree.Element, node_wg *sync.WaitGroup) {
	defer node_wg.Done()
	id := e.SelectAttrValue("id", "none")
	suspend_str := e.SelectAttrValue("suspend", "none")
	suspend, err := strconv.ParseInt(suspend_str, 10, 0)
	if err != nil {
		suspend = 0
	}
	var node common.NodeInterface = NewSuspendNode(id, int(suspend))
	mp_mutex.Lock()
	mp[id] = node
	mp_mutex.Unlock()
}
