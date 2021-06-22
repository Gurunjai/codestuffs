package list

import "fmt"

// Node - to hold the single linked list
type Node struct {
	val int
	next *Node
}

// DoubleNode - to hold the doubly linked list
type DoubleNode struct {
	val int
	next, prev *DoubleNode
}

// DoubleList Wrapper to double linked list
type DoubleList struct {
	head, tail *DoubleNode
}

// NewNode to provide the single linked list
func NewNode(val int) *Node {
	return &Node{
		val : val,
		next : nil,
	}
}

func newDoubleNode(val int) *DoubleNode {
	return &DoubleNode{
		val : val,
		prev : nil,
		next : nil,
	}
}

// NewDoubleList provides the doubly linked list
func NewDoubleList() *DoubleList {
	return &DoubleList{
		head : nil,
		tail : nil,
	}
}

// InsertAtHead - Insert a give value into the Head of the Doubly Linked List
func (dl *DoubleList) InsertAtHead(val int) {
	tmp := newDoubleNode(val)

	if dl.head == nil {
		dl.head, dl.tail = tmp, tmp
	} else {
		dl.tail.prev = dl.head.prev
		tmp.next = dl.head
		dl.head = tmp
	}
}

// InsertAtTail - Insert a give value into the Tail of the Doubly Linked List
func (dl *DoubleList) InsertAtTail(val int) {
	tmp := newDoubleNode(val)

	if dl.head == nil {
		dl.head = tmp
	} else {
		tmp.prev = dl.tail
		dl.tail.next = tmp
	}

	dl.tail = tmp
}

// Find - to find the element in the doubly linked list
func (dl *DoubleList) Find(val int) *DoubleNode {
	if dl.head.val == val {
		return dl.head
	}

	if dl.tail.val == val {
		return dl.tail
	}

	cur := dl.head
	for cur != nil {
		if cur.val == val {
			return cur
		}

		cur = cur.next
	}

	return nil
}

// Delete - to delete a value from the doubly linked list
func (dl *DoubleList) Delete(val int) {
	if dl.head.val == val {
		dl.head = dl.head.next
		return
	}

	if dl.tail.val == val {
		dl.tail.prev.next = dl.tail.next
		dl.tail = dl.tail.prev
		return
	}

	cur := dl.head
	for cur.next != nil {
		if cur.next.val == val {
			cur.next = cur.next.next
			cur.next.prev = cur
			return
		}

		cur = cur.next
	}

	return
}

// Reverse - Reverse a doubly linked list?? Is it really required?
func (dl *DoubleList) Reverse() *DoubleList {
	out := NewDoubleList()

	cur := dl.tail
	for cur != nil {
		out.InsertAtTail(cur.val)
		cur = cur.prev
	}

	return out
}

// Walk - walkthrough the doubly linked list
func (dl *DoubleList) Walk() {
	if dl.head == nil || dl.tail == nil {
		return
	}

	cur := dl.head
	for cur.next != nil {
		fmt.Printf(" %v, ", cur.val)
		cur = cur.next
	}

	fmt.Printf(" %v \n", cur.val)
}

// InsertAtHead - insert the val into the head of the list
func InsertAtHead(head *Node, val int) *Node {
	tmp := NewNode(val)
	if head != nil {
		tmp.next = head
	} 
	head = tmp

	return head
}

// InsertAtTail - insert the val into the tail of the list
func InsertAtTail(head *Node, val int) *Node {
	tmp := NewNode(val)
	if head == nil {
		head = tmp
	} else {
		curr := head
		for ; curr.next != nil; curr = curr.next {}
		curr.next = tmp
	}

	return head
}

// Find - find the first occurance of the node with the value provided
func Find(head *Node, val int) *Node {
	curr := head
	for ; curr != nil; curr = curr.next {
		if val == curr.val {
			return curr
		}
	}

	return nil
}

// Delete - delete the node with the provided value
func Delete(head *Node, val int) *Node {
	if head == nil {
		return head
	}

	if head.val == val {
		head = head.next
		return head
	}

	prev, curr := head, head
	for curr != nil {
		if curr.val == val {
			prev.next = curr.next
			return head
		}

		prev = curr
		curr = curr.next
	}

	return head
}

// DeleteSingle - delete the node with the provided value with single pointer
func DeleteSingle(head *Node, val int) *Node {
	if head == nil {
		return head
	}

	if head.val == val {
		head = head.next
		return head
	}

	curr := head
	for curr.next != nil {
		if curr.next.val == val {
			curr.next = curr.next.next
			return head
		}

		curr = curr.next
	}

	return head
}

// Reverse - reverse the order of the linked list
func Reverse(head *Node) *Node {
	var prev *Node
	curr := head
	for curr != nil {
		t := curr.next
		curr.next = prev
		prev = curr
		curr = t
	}

	return prev
}

// Walk - dump the list for all nodes
func Walk(head *Node) {
	tmp := head
	for tmp.next != nil {
		fmt.Printf(" %v, ", tmp.val)
		tmp = tmp.next
	}
	fmt.Printf(" %v\n", tmp.val)
}