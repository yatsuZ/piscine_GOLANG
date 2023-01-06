package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 10 && n >= 0 {
		z01.PrintRune(rune(n + 48))
	} else if n >= 10 {
		v := 0
		var L []int
		for n != 0 {
			v = n % 10
			n = n / 10
			L = append(L, v)
		}
		SortIntegerTable(L)
		for i := range L {
			z01.PrintRune(rune(L[i] + 48))
		}
	}
}
