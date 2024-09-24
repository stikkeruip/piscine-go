package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	curr := l
	temp := &NodeI{Data: data_ref}
	prev := temp

	if l.Data > data_ref {
		temp.Next = l
		return temp
	}

	for curr != nil {
		if curr.Data > data_ref {
			prev.Next = temp
			temp.Next = curr
			return l
		}
		prev = curr
		curr = curr.Next
	}

	prev.Next = temp
	return l
}
