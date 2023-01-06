package piscine

func ListReverse(l *List) {
	t := &List{}
	long := ListSize(l) - 1
	for i := long; i >= 0; i-- {
		ListPushBack(t, ListAt(l.Head, i).Data)
	}
	*l = *t
}
