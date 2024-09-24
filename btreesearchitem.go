package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if elem == root.Data {
		return root
	}
	if elem > root.Data {
		BTreeSearchItem(root.Right, elem).Parent = root
		return root.Right
	} else {
		BTreeSearchItem(root.Left, elem).Parent = root
		return root.Left
	}
}
