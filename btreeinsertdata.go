package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root.Data > data {
		if root.Left == nil {
			left := &TreeNode{Parent: root, Data: data}
			root.Left = left
			return root
		}
		return BTreeInsertData(root.Left, data)
	} else if root.Data < data {
		if root.Right == nil {
			right := &TreeNode{Parent: root, Data: data}
			root.Right = right
			return root
		}
		return BTreeInsertData(root.Right, data)
	}
	return BTreeInsertData(root.Right, data)
}
