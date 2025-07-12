package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: nil}
	prev := dummy
	current := head
	for current != nil {
		if current.Next == nil {
			prev.Next = current
			prev = prev.Next
		} else if current.Next.Val == current.Val {
			for current.Next != nil && current.Next.Val == current.Val {
				current = current.Next
			}
		} else {
			prev.Next = current
			prev = prev.Next
		}
		current = current.Next
	}
	prev.Next = nil
	return dummy.Next
}

func AddNode(head *ListNode, value int) *ListNode {
	newNode := &ListNode{Val: value}

	if head == nil {
		return newNode
	}

	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	return head
}

func PrintList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Val, " -> ")
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	var head *ListNode
	var head2 *ListNode
	nums := []int{1, 2, 3, 3, 4, 4, 5}
	for _, num := range nums {
		head = AddNode(head, num)
	}
	PrintList(head)
	newHead := deleteDuplicates(head)
	PrintList(newHead)
	nums2 := []int{1, 1, 1, 2, 3}
	for _, num := range nums2 {
		head2 = AddNode(head2, num)
	}
	PrintList(head2)
	newHead2 := deleteDuplicates(head2)
	PrintList(newHead2)
}
