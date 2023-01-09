package LZTree

import (
	"fmt"
)

// Tree
type Tree struct {
	nodes       []*Node
	sequence    []int
	root        *Node
	currentNode *Node
	totalBytes  int
}

func NewTree() *Tree {
	tree := &Tree{}
	tree.sequence = make([]int, 0)
	tree.nodes = make([]*Node, 0)
	tree.root = NewNode(-1, -2, nil, 0)
	tree.totalBytes = 0
	return tree
}

// добавить новый байт в последовательность
// если байт найден в дереве - вернуть -1 иначе - вернуть id созданного узла
func (t *Tree) Append(b byte) int {
	t.totalBytes = t.totalBytes + 1

	var node *Node

	// если есть пред. поиск - продолжаем поиск глубже
	if t.currentNode != nil {
		node = t.currentNode.FindTreeNode(b)
	} else {
		// иначе начинаем заново
		//--- find node
		node = t.root.FindTreeNode(b)
	}

	// если нода найдена - сохраняем как пред.
	if node != nil {
		t.currentNode = node
		return NO_NODE_ID
	}

	// не найдена
	return t.createNode(b)

}

// создать новый узел
func (t *Tree) createNode(data byte) int {
	node_id := len(t.nodes)

	var n *Node
	if t.currentNode != nil {
		n = NewNode(node_id, t.currentNode.ID, t.currentNode, data)
		t.currentNode.Append(n)
	} else {
		n = NewNode(node_id, -1, nil, data)
		t.root.Append(n)
	}

	t.currentNode = nil

	t.nodes = append(t.nodes, n)
	t.sequence = append(t.sequence, node_id)

	return node_id

}

// Finish - корректировка последней ноды
func (t *Tree) Finish() {
	if t.currentNode != nil {
		t.sequence = append(t.sequence, t.currentNode.ID)
		t.currentNode = nil
	}

}

func (t *Tree) PrinfInfo() {
	fmt.Println("========================================================")
	fmt.Println("total bytes count: ", t.totalBytes)
	fmt.Println("nodes count: ", len(t.nodes))
	fmt.Println("sequence count: ", len(t.sequence))
	// fmt.Println("dSaize\tcSize\tchunks\ttCount\ttotal\telapsed\t\tratio")
	// fmt.Printf("%d\t%d\t%d\t%d\t%d\t%s\t%d\n",
	// 	fileInfo.Size(), chunkSize, chunks, treeCount, totalCount, elapsed, ratio)

	maxLevel := 0
	for _, node := range t.nodes {
		if node.Level > maxLevel {
			maxLevel = node.Level
		}
	}

	fmt.Println("max branch level: ", maxLevel)

	fmt.Println("========================================================")
}

func (t *Tree) Unpack() {

	result := make([]byte, 0)

	for _, node_index := range t.sequence {
		node := t.nodes[node_index]
		// fmt.Println(node_index, node.Data)

		// // корневой узел - данные готовы
		// if node.Parent == nil {
		// 	fmt.Println(node.Data)
		// } else {
		// 	// иначе - собираем полный путь до корня

		// }

		// fmt.Println(node.GetData())
		result = append(result, node.GetData()...)
	}

	fmt.Println(result)

}

func (t *Tree) GetNodes() []*Node {
	return t.nodes
}

func (t *Tree) GetSequence() []int {
	return t.sequence
}

func (t *Tree) SetSequence(items []int) {
	t.sequence = items
}

func (t *Tree) SetNodes(nodes []*Node) {
	t.nodes = nodes
}
