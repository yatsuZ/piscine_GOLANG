package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return returnNbrBase(AtoiBase(nbr, baseFrom), baseTo)
}

func returnNbrBase(nbr int, base string) string {
	L := ""
	B := []rune(base)
	negatif := false
	reste := 0
	if nbr < 0 {
		negatif = true
		reste = (nbr % len(B)) * -1
		L = string(B[reste]) + L
		nbr = (nbr / len(B)) * -1
	}
	if nbr < len(B) {
		L = string(B[nbr]) + L
	} else {
		for nbr > 0 {
			reste = nbr % len(B)
			L = string(B[reste]) + L
			nbr = (nbr / len(B))
		}
	}
	if negatif {
		L = "-" + L
	}
	return L
}
