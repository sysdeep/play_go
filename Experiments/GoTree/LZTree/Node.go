package LZTree

// Node
type Node struct {
	ID        int
	data      byte
	childrens []*Node
}

func NewNode(id int, b byte) *Node {
	node := &Node{
		ID:        id,
		data:      b,
		childrens: make([]*Node, 0),
	}

	return node
}

func (n *Node) FindTreeNode(b byte) *Node {
	for _, node := range n.childrens {
		if node.data == b {
			return node
		}
	}

	return nil
}

func (n *Node) Append(node *Node) {
	n.childrens = append(n.childrens, node)
}
