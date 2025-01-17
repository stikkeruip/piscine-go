package piscine

func ListReverse(l *List) {
	curr := l.Head
	var prev *NodeL

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	temp := l.Head
	l.Head = prev
	l.Tail = temp
}
