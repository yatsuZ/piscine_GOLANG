package piscine

import (
	//	"fmt"

	"github.com/01-edu/z01"
)

func EightQueens() {
	//	combien := 0
	L := [8]int{1, 5, 8, 6, 3, 7, 2, 4}
	for L[0] != 0 {
		if verif(L) {
			//			combien++
			for _, v := range L {
				z01.PrintRune(rune(v + 48))
				//		fmt.Println(i, v)
			}
			z01.PrintRune('\n')
		}
		L = plus(L)
	}
	//	fmt.Println(combien)
}

func verif(L [8]int) bool {
	for i, v := range L {
		for j, m := range L {
			if i != j {
				if v-m == j-i {
					return false
				}
				if m-v == j-i {
					return false
				}
				if v == m {
					return false
				}
			}
		}
	}

	return true
}

func plus(L [8]int) [8]int {
	for i := range L {
		if L[7-i] != 8 {
			L[7-i] = 1 + L[7-i]
			break
		} else {
			if 7-i == 0 {
				return [8]int{0, 0, 0, 0, 0, 0, 0, 0}
			} else {
				L[7-i] = 1
			}
		}
	}
	return L
}
