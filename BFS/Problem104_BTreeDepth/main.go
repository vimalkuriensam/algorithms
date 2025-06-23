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

func maxDepth(root *TreeNode) int {
	var depth int
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		queueSize := len(queue)
		for range queueSize {
			node := Dequeue(&queue)
			if node.Left != nil {
				Enqueue(&queue, node.Left)
			}
			if node.Right != nil {
				Enqueue(&queue, node.Right)
			}
		}
		depth++
	}
	return depth
}
