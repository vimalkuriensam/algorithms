package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextNode(node **ListNode) {
	if (*node) != nil {
		(*node) = (*node).Next
	} else {
		(*node) = nil
	}
}

func nodeVal(node *ListNode) int {
	if node != nil {
		return node.Val
	}
	return 0
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	node := &ListNode{}
	current := node
	for l1 != nil || l2 != nil || carry != 0 {
		x := nodeVal(l1)
		y := nodeVal(l2)
		sum := x + y + carry
		carry = sum / 10
		new_node := &ListNode{Val: sum % 10}
		current.Next = new_node
		nextNode(&current)
		nextNode(&l1)
		nextNode(&l2)
	}
	return node.Next
}

func main() {}
