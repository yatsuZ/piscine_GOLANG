package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l.Next == nil {
		return l
	}
	t := l
	var liste []int
	for t != nil {
		liste = append(liste, t.Data)
		t = t.Next
	}
	SortIntegerTable(liste)
	machin := l
	for _, v := range liste {
		machin.Data = v
		machin = machin.Next
	}
	return l
}
