package argo

// The implementation of NodeInterface
type Node struct {
	id     string
	custom int
	in     []string
	out    []string
}

func (n *Node) GetId() string {
	return n.id
}

func (n *Node) GetCustom() int {
	return n.custom
}

func (n *Node) GetInNode() []string {
	return n.in
}

func (n *Node) GetOutNode() []string {
	return n.out
}

func (n *Node) AppendInNode(s string) {
	n.in = append(n.in, s)
}

func (n *Node) AppendOutNode(s string) {
	n.out = append(n.out, s)
}

func (n *Node) HaveInNode() bool {
	if len(n.in) > 0 {
		return true
	} else {
		return false
	}
}

func (n *Node) HaveOutNode() bool {
	if len(n.out) > 0 {
		return true
	} else {
		return false
	}
}
