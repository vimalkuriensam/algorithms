package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)

	if leftDepth == 0 {
		return 1 + rightDepth
	} else if rightDepth == 0 {
		return 1 + leftDepth
	}
	return 1 + min(leftDepth, rightDepth)
}

func main() {}
