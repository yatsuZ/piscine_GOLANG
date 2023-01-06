package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}
	t := l
	noeud := t.Head
	ListClear(l)
	for noeud != nil {
		if noeud.Data != data_ref {
			ListPushBack(l, noeud.Data)
		}
		noeud = noeud.Next
	}
}
