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

func levelOrder(root *TreeNode) [][]int {
	var lOrder [][]int
	if root == nil {
		return lOrder
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		queue_size := len(queue)
		var current_queue []int
		for i := 0; i < queue_size; i++ {
			current_node := dequeue(&queue)
			if current_node != nil {
				current_queue = append(current_queue, current_node.Val)
				if current_node.Left != nil {
					enqueue(&queue, current_node.Left)
				}
				if current_node.Right != nil {
					enqueue(&queue, current_node.Right)
				}
			}
		}
		lOrder = append(lOrder, current_queue)
	}
	return lOrder
}

func main() {

}
