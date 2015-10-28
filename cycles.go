package linkedListRiddles

// You get a node that is part of a linked list. The node has a next element and data.
// Your goal is to detect if the linked list contains cycles! Furthermore you must
// NOT alter the list in any way, your memory usage should be INDEPENDENT of the list's
// length (so no keeping track of all visited nodes) and should run in O(n)... so linear time

// The idea is simple: we loop and move the head forwards by 1 every time. if we reach nil,
// there's no loop. we drag behnd a second reference, which we move forwards every second time
// if we have a loop eventually the dragged behind value and the head are the same: we detected
// a loop. memory is clearly constant. we did not alter the list. worst time szenario is one big loop
// where we have time consumption 2 * n, if n is the length of the list, which is still in O(n).
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
