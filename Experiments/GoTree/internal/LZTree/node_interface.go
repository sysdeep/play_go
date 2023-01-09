package LZTree

type NodeInterface interface {
	AppendChildren(NodeInterface)
}
