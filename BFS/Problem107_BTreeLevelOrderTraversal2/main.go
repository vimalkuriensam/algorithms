package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func enqueu(queue *[]*TreeNode, node *TreeNode) {
	(*queue) = append((*queue), node)
}

func dequeue(queue *[]*TreeNode) *TreeNode {
	node := (*queue)[0]
	(*queue) = (*queue)[1:]
	return node
}

func reverse_level_order(level_order [][]int) {
	for i, j := 0, len(level_order)-1; i < j; i, j = i+1, j-1 {
		level_order[i], level_order[j] = level_order[j], level_order[i]
	}
}

func levelOrderBottom(root *TreeNode) [][]int {
	var levelOrder [][]int
	if root == nil {
		return levelOrder
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		queue_size := len(queue)
		var inner_queue []int
		for i := 0; i < queue_size; i++ {
			current_node := dequeue(&queue)
			if current_node != nil {
				inner_queue = append(inner_queue, current_node.Val)
				if current_node.Left != nil {
					enqueu(&queue, current_node.Left)
				}
				if current_node.Right != nil {
					enqueu(&queue, current_node.Right)
				}
			}
		}
		levelOrder = append(levelOrder, inner_queue)
	}
	reverse_level_order(levelOrder)
	return levelOrder
}

func main() {}
