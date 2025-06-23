package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Enqueue(queue *[]*TreeNode, node *TreeNode) {
	(*queue) = append((*queue), node)
}

func Dequeue(queue *[]*TreeNode) *TreeNode {
	node := (*queue)[0]
	(*queue) = (*queue)[1:]
	return node
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left_queue := []*TreeNode{root.Left}
	right_queue := []*TreeNode{root.Right}
	for len(left_queue) > 0 && len(right_queue) > 0 {
		left_node := Dequeue(&left_queue)
		right_node := Dequeue(&right_queue)
		if left_node == nil && right_node == nil {
			continue
		}
		if (left_node == nil) != (right_node == nil) {
			return false
		}
		if left_node.Val != right_node.Val {
			return false
		}
		Enqueue(&left_queue, left_node.Left)
		Enqueue(&left_queue, left_node.Right)
		Enqueue(&right_queue, right_node.Right)
		Enqueue(&right_queue, right_node.Left)
	}
	return true
}
