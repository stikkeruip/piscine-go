package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	curr := l

	for i := 0; i < pos; i++ {
		if curr == nil {
			return nil
		}

		curr = curr.Next
	}
	return curr
}
