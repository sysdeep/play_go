package LZTree

import "testing"

// Success and failure markers.
const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestCreate(t *testing.T) {
	t.Log("Create new node")

	node := NewNode(0, 1, nil, byte(0xff))

	if node.Level != 0 {
		// t.Error("Expected 0, got: ", node.Level)
		t.Fatalf("\t%s\tExpected 0, got: %d", failed, node.Level)
	}

	t.Logf("\t%s\tNode Level is 0", success)
}

func TestAppendNode(t *testing.T) {
	t.Log("Append node")

	root := NewNode(0, 1, nil, byte(0xff))

	node := NewNode(1, 2, root, byte(0x11))
	root.Append(node)

	if node.Level == 0 {
		t.Fatalf("Expected 1, got %d", node.Level)
	}

}
