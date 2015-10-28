package linkedListRiddles

import (
	"testing"
)

func TestCycles(t *testing.T) {
	var n1, n2, n3, n4, n5 Node

	if HasCycle(n1) {
		t.Error("A list with a single element does not have cycles...")
	}

	n1.next = &n2
	if HasCycle(n1) {
		t.Error("n1 -> n2 -> nil. Does not have cylcles.")
	}

	n1.next = &n2
	n2.next = &n1
	if !HasCycle(n1) {
		t.Error("n1 -> n2 -> n1. Does have cylces.")
	}

	n1.next = &n2
	n2.next = &n3
	n3.next = &n4
	n4.next = &n5
	n5.next = &n3
	if !HasCycle(n1) {
		t.Error("n1 -> n2 -> n3 -> n4 -> n5 -> n3. Does have cylces.")
	}

	n5.next = nil

	if HasCycle(n1) {
		t.Error("n1-> n2->n3->n4->n5->nil does not have cycles")
	}

}
