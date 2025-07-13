package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	dummy := &ListNode{Next: head}
	listLength := 0
	current := dummy
	for current.Next != nil {
		listLength++
		current = current.Next
	}
	current.Next = head
	current = dummy.Next
	cutoffLength := listLength - (k % listLength)
	for i := 0; i < cutoffLength-1; i++ {
		current = current.Next
	}
	newHead := current.Next
	current.Next = nil
	return newHead
}
