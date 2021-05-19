package listnode

import (
	// "fmt"
	"testing"
)

func TestSLLInsertAtHead(t *testing.T) {
	in := []int{10, 20, 30, 40, 50, 60, 70}

	var head *Node;
	for _, v := range(in) {
		head = InsertAtHead(head, v)
	}

	tmp := head
	for i := len(in) - 1; i >= 0; i-- {
		if tmp.val != in[i] {
			t.Fatalf("Mismatch detected:\n \t\t Node Value: %v, \t Input Value: %v\n", tmp.val, in[i])
		}

		tmp = tmp.next
	}
}

func TestSLLInsertAtTail(t *testing.T) {
	in := []int{10, 20, 30, 40, 50, 60, 70}

	var head *Node;
	for _, v := range(in) {
		head = InsertAtTail(head, v)
	}

	tmp := head
	for i := 0; i < len(in); i++ {
		if tmp.val != in[i] {
			t.Fatalf("Mismatch detected:\n \t\t Node Value: %v, \t Input Value: %v\n", tmp.val, in[i])
		}

		tmp = tmp.next
	}
}

func TestSLLInsertAndFind(t *testing.T) {
	in := []int{10, 20, 30, 40, 50, 60, 70}
	const find int = 60

	var head *Node;
	for _, v := range(in) {
		head = InsertAtTail(head, v)
	}

	tmp := Find(head, find)
	
	if tmp != nil && tmp.val != find {
		t.Fatalf("Fetched a wrong value(%v) instead of %v\n", tmp.val, find)
	}
}

func TestSLLInsertDeleteWalk(t *testing.T) {
	in := []int{10, 20, 30, 40, 50, 60, 70}
	const del int = 10

	var head *Node;
	for _, v := range(in) {
		head = InsertAtTail(head, v)
	}

	head = Delete(head, del)
	
	if Find(head, del) != nil {
		t.Fatalf("Deleted value is present in the node\n")
	}
}

func TestSLLInsertDeleteSingleWalk(t *testing.T) {
	in := []int{10, 20, 30, 40, 50, 60, 70}
	const del int = 10

	var head *Node;
	for _, v := range(in) {
		head = InsertAtTail(head, v)
	}

	head = DeleteSingle(head, del)

	if Find(head, del) != nil {
		t.Fatalf("Deleted value is present in the node\n")
	}
}


func TestSLLInsertReverseWalk(t *testing.T) {
	in := []int{10, 20, 30, 40, 50, 60, 70}

	var head *Node;
	for _, v := range(in) {
		head = InsertAtTail(head, v)
	}

	tmp := Reverse(head)
	for i := len(in) - 1; i >= 0; i-- {
		if tmp.val != in[i] {
			t.Fatalf("Mismatch detected:\n \t\t Node Value: %v, \t Reverse Input Value: %v\n", tmp.val, in[i])
		}

		tmp = tmp.next
	}
}

func TestDLLInsertAtHead(t *testing.T) {
	in := []int{10, 20, 30, 40, 50, 60, 70, 80}

	head := newDoubleList()
	for _, v := range(in) {
		head.InsertAtHead(v)
	}

	cur := head.head
	for i := len(in) - 1; i >= 0; i-- {
		if cur.val != in[i] {
			t.Fatalf("Mismatch found:\n \t\t Node Value: %v, \t Input Value: %v\n", cur.val, in[i])
		}
		cur = cur.next
	}	
}

func TestDLLInsertAtTail(t *testing.T) {
	in := []int{10, 20, 30, 40, 50, 60, 70, 80}

	head := newDoubleList()
	for _, v := range(in) {
		head.InsertAtTail(v)
	}

	cur := head.head
	for i := 0; i < len(in); i++ {
		if cur.val != in[i] {
			t.Fatalf("Mismatch found:\n \t\t Node Value: %v, \t Input Value: %v\n", cur.val, in[i])
		}

		cur = cur.next
	}
}

func TestDLLInsertFind(t *testing.T) {
	in := []int{10, 20, 30, 40, 50, 60, 70, 80}
	const find int = 30

	head := newDoubleList()
	for _, v := range(in) {
		head.InsertAtTail(v)
	}

	tmp := head.Find(find)
	if tmp != nil && tmp.val != find {
		t.Fatalf("Fetched a wrong value(%v) instead of %v\n", tmp.val, find)
	}
}

func TestDLLInsertDelete(t *testing.T) {
	in := []int{10, 20, 30, 40, 50, 60, 70, 80}
	const del int = 30

	head := newDoubleList()
	for _, v := range(in) {
		head.InsertAtTail(v)
	}

	head.Delete(del)
	if head.Find(del) != nil {
		head.Walk()
		t.Fatalf("Deleted value is present in the node\n")
	}	
}

func TestDLLInsertReverse(t *testing.T) {
	in := []int{10, 20, 30, 40, 50, 60, 70, 80}
	
	head := newDoubleList()
	for _, v := range(in) {
		head.InsertAtTail(v)
	}

	ot := head.Reverse()
	cur := ot.head
	for i := len(in) - 1; i >= 0; i-- {
		if cur.val != in[i] {
			t.Fatalf("Mismatch Found:\n \t\t Node Value: %v, \t Reversed Input Value: %v\n", cur.val, in[i])
		}

		cur = cur.next
	}
}