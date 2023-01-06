package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) < 2 {
		z01.PrintRune(10)
	} else {
		var (
			Pmot []int
			Plet []int
			nbrV int
		)
		L := os.Args[1:]
		Pmot, Plet, nbrV = Vbool(L)
		if nbrV > 1 {
			count := 0
			impaire := -1
			if nbrV%2 == 1 {
				impaire = nbrV / 2
			}
			nbrV = nbrV - 1
			for mots := range L {
				for lettre := range L[mots] {
					if count <= nbrV {
						if mots == Pmot[count] && lettre == Plet[count] {
							if count != impaire {
								z01.PrintRune(rune(L[Pmot[nbrV-count]][Plet[nbrV-count]]))
							} else {
								z01.PrintRune(rune(L[mots][lettre]))
							}
							count++
						} else {
							z01.PrintRune(rune(L[mots][lettre]))
						}
					} else {
						z01.PrintRune(rune(L[mots][lettre]))
					}
				}
				z01.PrintRune(' ')
			}
			z01.PrintRune(10)
		} else {
			for i := range L {
				for j := range L[i] {
					z01.PrintRune(rune(L[i][j]))
				}
				z01.PrintRune(' ')
			}
			z01.PrintRune(10)
		}
	}
}

func Vbool(L []string) ([]int, []int, int) {
	nbrDEVoyelle := 0
	var PositionMot []int
	var PositionLettre []int
	for mots := range L {
		for lettre := range L[mots] {
			if rune(L[mots][lettre]) == 'A' || rune(L[mots][lettre]) == 'E' || rune(L[mots][lettre]) == 'I' || rune(L[mots][lettre]) == 'O' || rune(L[mots][lettre]) == 'U' || rune(L[mots][lettre]) == 'a' || rune(L[mots][lettre]) == 'e' || rune(L[mots][lettre]) == 'i' || rune(L[mots][lettre]) == 'o' || rune(L[mots][lettre]) == 'u' {
				nbrDEVoyelle++
				PositionMot = append(PositionMot, mots)
				PositionLettre = append(PositionLettre, lettre)
			}
		}
	}
	return PositionMot, PositionLettre, nbrDEVoyelle
}
