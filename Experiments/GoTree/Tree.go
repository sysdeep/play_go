package main

type Tree struct {
	childrens []Node
}

func NewTree() *Tree {
	t := &Tree{}

	return t
}

func (t *Tree) AppendChunk(chunk []byte) {

	// fmt.Println(chunk)
	// b0 b1 b2 b3
	for _, b := range chunk {
		// fmt.Println(b)
		node := t.GetOrCreate(b)
		node.AppendChunk(chunk[1:])
	}

}

func (t *Tree) GetOrCreate(b byte) *Node {
	for _, node := range t.childrens {
		if node.data == b {
			return &node
		}
	}

	node := NewNode(b)
	t.childrens = append(t.childrens, node)
	return &node
}

func (t *Tree) GetCount() int {
	result := len(t.childrens)
	for _, node := range t.childrens {
		result += node.GetCount()
	}
	return result
}
