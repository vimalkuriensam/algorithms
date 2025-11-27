package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftdepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return 1 + max(leftdepth, rightDepth)
}

func main() {}
