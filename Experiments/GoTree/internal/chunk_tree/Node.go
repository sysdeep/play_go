package chunk_tree

type Node struct {
	data         byte
	childrens    []Node
	childrensMap map[byte]Node
}

func NewNode(data byte) Node {
	n := Node{
		data:      data,
		childrens: []Node{},
		// childrensMap: make(map[byte]Node),
	}
	n.childrensMap = make(map[byte]Node)
	return n
}

func (n *Node) AppendChunk(chunk []byte) {

	//b1 b2 b3
	for _, b := range chunk {
		// log.Println(i, b)
		node := n.GetOrCreate(b)
		node.AppendChunk(chunk[1:])
	}

}

func (n *Node) GetOrCreate(b byte) *Node {

	//--- map
	// node, ok := n.childrensMap[b]
	// if ok {
	// 	// fmt.Println("get node")
	// 	return &node
	// }

	// // fmt.Println("create node for ", n)
	// // fmt.Println(n)
	// newNode := NewNode(b)
	// // fmt.Println(newNode)
	// // fmt.Println(n.childrensMap)
	// n.childrensMap[b] = newNode
	// // fmt.Println(n)
	// // n.childrens = append(n.childrens, node)
	// return &newNode

	//--- childrens
	for _, node := range n.childrens {
		if node.data == b {
			return &node
		}
	}

	newNode := NewNode(b)
	n.childrens = append(n.childrens, newNode)
	return &newNode

}

func (n *Node) GetCount() int {
	result := len(n.childrens)
	for _, node := range n.childrens {
		result += node.GetCount()
	}
	return result
}
