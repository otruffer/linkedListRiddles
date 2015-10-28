package linkedListRiddles

// You get a node that is part of a linked list. The node has a next element and data.
// Your goal is to detect if the linked list contains cycles! Furthermore you must
// NOT alter the list in any way, your memory usage should be INDEPENDENT of the list's
// length (so no keeping track of all visited nodes) and should run in O(n)... so linear time
func HasCycle(head *Node) bool {
	var drag = head
	for i := 0; head.hasNext(); i++ {
		if i%2 == 1 {
			drag = drag.next
		}
		head = head.next
		if drag == head {
			return true
		}
	}

	return false
}
