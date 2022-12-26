package parser

import (
	"encoding/json"

	"github.com/beevik/etree"
)

const (
	DockerType = "1"
)

type Node interface {
	// Obtain basic information about the node
	GetName() string
	GetDependencies() []string
	GetCustom() string
	AppendDep(string)

	// Each node parses its own information
	Fillin(*etree.Element) error
}

type NodeSet map[string]Node

func (s *NodeSet) Marshal() ([]byte, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return b, nil
}

type BaseNode struct {
	Name         string
	Dependencies []string
	Custom       string
}

func (n *BaseNode) GetName() string {
	return n.Name
}

func (n *BaseNode) GetDependencies() []string {
	return n.Dependencies
}

func (n *BaseNode) GetCustom() string {
	return n.Custom
}

func (n *BaseNode) AppendDep(node string) {
	n.Dependencies = append(n.Dependencies, node)
}

type DockerNode struct {
	BaseNode
	Action string
	URL    string
}

func NewDockerNode() *DockerNode {
	return &DockerNode{
		BaseNode: BaseNode{
			Dependencies: make([]string, 0),
			Custom:       DockerType,
		},
	}
}

func (d *DockerNode) Fillin(element *etree.Element) error {
	// TODO: Parse element to node information
	return nil
}
