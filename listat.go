package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	for i := 0; i < pos; i++ {
		if l == nil {
			return nil
		}
		l = l.Next
	}
	return l
}
