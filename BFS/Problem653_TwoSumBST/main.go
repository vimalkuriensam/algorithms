package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func enqueue(queue *[]*TreeNode, node *TreeNode) {
	*(queue) = append(*(queue), node)
}

func dequeue(queue *[]*TreeNode) *TreeNode {
	node := (*queue)[0]
	*queue = (*queue)[1:]
	return node
}

func findTarget(root *TreeNode, k int) bool {
	queue := []*TreeNode{root}
	hashMap := make(map[int]struct{})
	for len(queue) > 0 {
		currentNode := dequeue(&queue)
		if currentNode != nil {
			complement := k - currentNode.Val
			if _, found := hashMap[complement]; found {
				return true
			}
			hashMap[currentNode.Val] = struct{}{}
			enqueue(&queue, currentNode.Left)
			enqueue(&queue, currentNode.Right)
		}
	}
	return false
}

func main() {}
