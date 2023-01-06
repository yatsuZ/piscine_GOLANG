package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	L := ""
	B := []rune(base)
	negatif := false
	if bonneBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
	} else {
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
		for _, i := range L {
			z01.PrintRune(rune(i))
		}
	}
}

func bonneBase(base string) bool {
	L := []rune(base)
	if len(L) < 2 {
		return true
	}
	for i := 0; i < len(L); i++ {
		for j := 0; j < len(L); j++ {
			if L[i] == L[j] && i != j || L[i] == '-' || L[i] == '+' || L[j] == '-' || L[j] == '+' {
				return true
			}
		}
	}
	return false
}
