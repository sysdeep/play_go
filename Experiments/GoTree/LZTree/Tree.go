package LZTree

import (
	"fmt"
	"log"
)

// Tree
type Tree struct {
	nodes    []*Node
	sequence []int
	root     *Node
	last     *Node
	total    int
}

func NewTree() *Tree {
	tree := &Tree{}
	tree.sequence = make([]int, 0)
	tree.nodes = make([]*Node, 0)
	tree.root = NewNode(-1, 0)
	tree.total = 0
	return tree
}

func (t *Tree) Append(b byte) {
	t.total = t.total + 1
	var node *Node

	// если есть пред. поиск - продолжаем поиск глубже
	if t.last != nil {
		node = t.last.FindTreeNode(b)
	} else {
		// иначе начинаем заново
		//--- find node
		node = t.root.FindTreeNode(b)
	}

	if node != nil {
		// если нода найдена - сохраняем как пред.
		log.Println("Tree: node found")
		// n := NewNode(b)
		// node.Append(n)
		// t.nodes = append(t.nodes, n)
		t.last = node
	} else {
		// не найдена

		node_id := len(t.nodes)

		log.Println("Tree: node NOT found - create: ", node_id)
		n := NewNode(node_id, b)

		if t.last != nil {
			t.last.Append(n)
		} else {
			t.root.Append(n)
		}

		t.last = nil
		// t.root.Append(n)
		t.nodes = append(t.nodes, n)
		t.sequence = append(t.sequence, node_id)
	}

}

// Finish - корректировка последней ноды
func (t *Tree) Finish() {
	if t.last != nil {
		t.sequence = append(t.sequence, t.last.ID)
		t.last = nil
	}

}

func (t *Tree) PrinfInfo() {
	fmt.Println("========================================================")
	fmt.Println("total bytes count: ", t.total)
	fmt.Println("nodes count: ", len(t.nodes))
	fmt.Println("sequence count: ", len(t.sequence))
	// fmt.Println("dSaize\tcSize\tchunks\ttCount\ttotal\telapsed\t\tratio")
	// fmt.Printf("%d\t%d\t%d\t%d\t%d\t%s\t%d\n",
	// 	fileInfo.Size(), chunkSize, chunks, treeCount, totalCount, elapsed, ratio)

	fmt.Println("========================================================")
}

// func (t *Tree) findNode(b byte) *Node {
// 	for _, node := range t.nodes {
// 		if node.data == b {
// 			return node
// 		}
// 	}

// 	return nil
// }
