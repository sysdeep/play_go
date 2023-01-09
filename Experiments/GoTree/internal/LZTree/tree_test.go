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

func TestAppendOneByte(t *testing.T) {
	t.Log("Test append one byte")

	tree := NewTree()
	node_id := tree.Append(0x00)

	if node_id == -1 {
		t.Fatalf("Expected create node and return its id, got %d", node_id)
	}

}

func TestAppendTwoSameBytes(t *testing.T) {
	t.Log("Test append two same byts")
	byte_to_add := 0x00
	tree := NewTree()
	node_id := tree.Append(byte(byte_to_add))

	if node_id == -1 {
		t.Fatalf("Expected create node and return its id, got %d", node_id)
	}

	node_id = tree.Append(byte(byte_to_add))

	if node_id != -1 {
		t.Fatalf("Expected use same node and return no id, got %d", node_id)
	}

}
