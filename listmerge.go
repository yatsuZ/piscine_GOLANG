package piscine

func ListMerge(l1 *List, l2 *List) {
	t := l2.Head
	for t != nil {
		ListPushBack(l1, t.Data)
		t = t.Next
	}
}
