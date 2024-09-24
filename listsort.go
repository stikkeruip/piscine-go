package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	var mid *NodeI
	var prev *NodeI
	fast := l
	slow := l

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	mid = slow

	left := ListSort(l)
	right := ListSort(mid)

	return mergeLists(left, right)
}

func mergeLists(l1, l2 *NodeI) *NodeI {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Data < l2.Data {
		l1.Next = mergeLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeLists(l1, l2.Next)
		return l2
	}
}
