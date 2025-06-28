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

func zigzagLevelOrder(root *TreeNode) [][]int {
	var levelOrder [][]int
	if root == nil {
		return levelOrder
	}
	queue := []*TreeNode{root}
	var zigzag bool
	for len(queue) > 0 {
		queue_size := len(queue)
		var current_queue []int
		for i := 0; i < queue_size; i++ {
			current_node := dequeue(&queue)
			if current_node != nil {
				if current_node.Left != nil {
					enqueue(&queue, current_node.Left)
				}
				if current_node.Right != nil {
					enqueue(&queue, current_node.Right)
				}
				if !zigzag {
					current_queue = append(current_queue, current_node.Val)
				} else {
					current_queue = append([]int{current_node.Val}, current_queue...)
				}
			}
		}
		levelOrder = append(levelOrder, current_queue)
		zigzag = !zigzag
	}
	return levelOrder
}

func main() {}
