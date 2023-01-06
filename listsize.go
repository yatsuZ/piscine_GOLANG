package piscine

func ListSize(l *List) int {
	count := 0
	n := l.Head
	for n != nil {
		count++
		n = n.Next
	}
	return count
}
