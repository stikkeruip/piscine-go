package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Parent != nil {
		if root.Data < root.Parent.Data {
			if root == root.Parent.Left {
				return true
			}
			return false
		} else {
			if root == root.Parent.Right {
				return true
			}
			return false
		}
	}

	left := BTreeIsBinary(root.Left)
	right := BTreeIsBinary(root.Right)

	return left && right
}
