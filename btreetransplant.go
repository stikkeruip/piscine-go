package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	node.Data = rplc.Data
	node.Left = rplc.Left
	node.Right = rplc.Right

	return root
}
