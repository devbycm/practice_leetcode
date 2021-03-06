package practice_leetcode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		left := root.Val - p.Val
		right := root.Val - q.Val
		if left > 0 && right > 0 {
			root = root.Left
		} else if left < 0 && right < 0 {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
