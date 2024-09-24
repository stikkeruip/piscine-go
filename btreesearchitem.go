package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if elem == root.Data {
		return root
	}
	if elem < root.Data {
		res := BTreeSearchItem(root.Left, elem)
		if res != nil {
			res.Parent = root
		}
		return res
	} else {
		res := BTreeSearchItem(root.Right, elem)
		if res != nil {
			res.Parent = root
		}
		return res
	}
}
