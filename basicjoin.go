package piscine

func BasicJoin(elems []string) string {
	if len(elems) <= 1 {
		return string(elems[0])
	} else {
		return BasicJoin(elems[:len(elems)/2]) + BasicJoin(elems[len(elems)/2:])
	}
}
