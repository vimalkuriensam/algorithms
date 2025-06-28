package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func enqueue(queue *[]*TreeNode, node *TreeNode) {
	(*queue) = append((*queue), node)
}

func dequeue(queue *[]*TreeNode) *TreeNode {
	node := (*queue)[0]
	(*queue) = (*queue)[1:]
	return node
}

func minDepth(root *TreeNode) int {
	var depth int
	if root == nil {
		return depth
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		queue_size := len(queue)
		depth++
		for i := 0; i < queue_size; i++ {
			current_node := dequeue(&queue)
			if current_node != nil {
				if current_node.Left == nil && current_node.Right == nil {
					return depth
				}
				if current_node.Left != nil {
					enqueue(&queue, current_node.Left)
				}
				if current_node.Right != nil {
					enqueue(&queue, current_node.Right)
				}
			}
		}
	}
	return depth
}

func main() {}
