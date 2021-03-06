package practice_leetcode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	a := maxDepth(root.Left)
	b := maxDepth(root.Right)
	if a > b {
		return a + 1
	} else {
		return b + 1
	}
}
