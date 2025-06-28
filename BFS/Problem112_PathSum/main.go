package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type QueueItem struct {
	node *TreeNode
	sum  int
}

func enqueue(queue *[]*QueueItem, node *QueueItem) {
	(*queue) = append((*queue), node)
}

func dequeue(queue *[]*QueueItem) *QueueItem {
	node := (*queue)[0]
	(*queue) = (*queue)[1:]
	return node
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	queue := []*QueueItem{{node: root, sum: root.Val}}
	for len(queue) > 0 {
		current_node := dequeue(&queue)
		if current_node != nil {
			if current_node.node.Left == nil && current_node.node.Right == nil {
				if current_node.sum == targetSum {
					return true
				}
			}
			if current_node.node.Left != nil {
				node := &QueueItem{node: current_node.node.Left, sum: current_node.sum + current_node.node.Left.Val}
				enqueue(&queue, node)
			}
			if current_node.node.Right != nil {
				node := &QueueItem{node: current_node.node.Right, sum: current_node.sum + current_node.node.Right.Val}
				enqueue(&queue, node)
			}
		}
	}
	return false
}

func main() {}
