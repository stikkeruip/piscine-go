package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	node.Data = rplc.Data

	return root
}
