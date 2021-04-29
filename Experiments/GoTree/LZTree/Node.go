package LZTree

// Node
type Node struct {
	ID        int
	ParentID  int
	Parent    *Node
	Data      byte
	childrens []*Node
}

func NewNode(id int, parentID int, parent *Node, b byte) *Node {
	node := &Node{
		ID:        id,
		ParentID:  parentID,
		Parent:    parent,
		Data:      b,
		childrens: make([]*Node, 0),
	}

	return node
}

func (n *Node) FindTreeNode(b byte) *Node {
	for _, node := range n.childrens {
		if node.Data == b {
			return node
		}
	}

	return nil
}

func (n *Node) Append(node *Node) {
	n.childrens = append(n.childrens, node)
}

func (n *Node) GetData() []byte {

	if n.Parent == nil {
		// нет родителя - только свои данные
		return []byte{n.Data}
	} else {
		// есть родитель - данные родителя и свои
		parentData := n.Parent.GetData()

		result := make([]byte, 0)
		result = append(result, parentData...)
		result = append(result, n.Data)
		return result
	}
}
