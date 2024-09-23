package piscine

func ListMerge(l1 *List, l2 *List) {
	curr1 := l1.Tail
	curr2 := l2.Head

	for curr2 != nil {
		curr1.Next = curr2
		if curr1.Next != nil {
			curr1 = curr1.Next
		}
		curr2 = curr2.Next
	}
	l1.Tail = curr1
}
