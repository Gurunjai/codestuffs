/**
 * Definition for singly-linked list.
 */
 type ListNode struct {
    Val int
    Next *ListNode
 }
 
 func NodeLen(node *ListNode) (cnt int) {
    for ; node != nil; node = node.Next {
        cnt += 1
    }
    
    return
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (out *ListNode) {
    if NodeLen(l2) > NodeLen(l1) {
        l1, l2 = l2, l1
    }
        
    var prevNode *ListNode
    var carry int
    var l2Val int
    out = l1
    for l1 != nil {
        if l2 != nil {
            l2Val = l2.Val
            l2 = l2.Next
        } else {
            l2Val = 0
        }
        l1.Val += l2Val + carry
        carry = l1.Val / 10
        l1.Val %= 10
        prevNode = l1
        l1 = l1.Next
    }
    
    if carry == 0 {
        return
    }
    
    node := new (ListNode)
    node.Val = carry
    prevNode.Next = node
    return
}