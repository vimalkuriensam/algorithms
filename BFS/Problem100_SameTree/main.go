package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SameTree(p *TreeNode, q *TreeNode) bool {
	queue1 := []*TreeNode{p}
	queue2 := []*TreeNode{q}
	for len(queue1) > 0 && len(queue2) > 0 {
		current_node_1 := Dequeu(&queue1)
		current_node_2 := Dequeu(&queue2)
		if current_node_1 == nil && current_node_2 == nil {
			continue
		}
		if current_node_1 == nil || current_node_2 == nil {
			return false
		}
		if current_node_1.Val != current_node_2.Val {
			return false
		}
		Enqueu(&queue1, current_node_1.Left)
		Enqueu(&queue1, current_node_1.Right)
		Enqueu(&queue2, current_node_2.Left)
		Enqueu(&queue2, current_node_2.Right)
	}
	return true
}

func Enqueu(queue *[]*TreeNode, node *TreeNode) {
	*queue = append(*queue, node)
}

func Dequeu(queue *[]*TreeNode) *TreeNode {
	node := (*queue)[0]
	*queue = (*queue)[1:]
	return node
}
