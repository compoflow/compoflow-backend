package parser

type Node struct {
	Name         string
	Action       string
	URL          string
	Dependencies []string
	Custom       string
}

func (n *Node) AppendDep(node string) {
	n.Dependencies = append(n.Dependencies, node)
}
