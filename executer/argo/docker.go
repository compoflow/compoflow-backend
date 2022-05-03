package argo

import (
	"sync"

	"github.com/Lavender-QAQ/microservice-workflows-backend/executer/common"
	"github.com/beevik/etree"
)

// Docker node
type DockerNode struct {
	Node
	image   string
	port    int
	target  string
	command []string
	args    []string
}

func NewDockerNode(id string, image string) *DockerNode {
	return &DockerNode{
		Node: Node{
			id:     id,
			custom: 1,
			in:     []string{},
			out:    []string{},
		},
		image: image,
	}
}

// Add docker node to map
func buildDockerNode(e etree.Element, node_wg *sync.WaitGroup) {
	defer node_wg.Done()
	id := e.SelectAttrValue("id", "none")
	image := e.SelectAttrValue("image", "none")
	var node common.NodeInterface = NewDockerNode(id, image)
	mp_mutex.Lock()
	mp[id] = node
	mp_mutex.Unlock()
}
