package LZTree

import "testing"

func TestAppend(t *testing.T) {
	t.Log("Test append data to tree")
	tree := NewTree()
	tree.Append(0x00)

	nodes := tree.GetNodes()
	if len(nodes) != 1 {
		t.Fatalf("Expected nodes count 1, got %d", len(nodes))
	}
}
