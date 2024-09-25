package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 1
	}

	return 1 + BTreeLevelCount(root.Left)
	return 1 + BTreeLevelCount(root.Right)
}
