package linkedListRiddles

type Node struct {
	data int
	next *Node
}

func (n *Node) hasNext() bool {
	return n.next != nil
}
