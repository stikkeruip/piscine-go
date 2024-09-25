package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	// Case 1: Node has two children
	if node.Left != nil && node.Right != nil {
		// Find the in-order successor (smallest node in the right subtree)
		successor := findMin(node.Right)

		// Replace the current node's data with the successor's data
		node.Data = successor.Data

		// Now "delete" the successor, but successor can only have one child (right child)
		node = successor
	}

	// Case 2 & 3: Node has one child or no children
	var child *TreeNode
	if node.Left != nil {
		child = node.Left
	} else {
		child = node.Right
	}

	// If node is the root
	if node.Parent == nil {
		root = child
		if child != nil {
			child.Parent = nil
		}
		return root
	}

	// Update parent's child pointer
	if node == node.Parent.Left {
		node.Parent.Left = child
	} else {
		node.Parent.Right = child
	}

	// Update child's parent pointer if child exists
	if child != nil {
		child.Parent = node.Parent
	}

	return root
}

// Helper function to find the minimum value node (in-order successor)
func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
