package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	noeud := l.Head
	for noeud != nil {
		if comp(noeud.Data, ref) {
			return &noeud.Data
		}
		noeud = noeud.Next
	}
	return nil
}
