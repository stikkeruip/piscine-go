package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		return nil
	}
	head := n1
	curr := n1

	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = n2

	return ListSort(head)
}
