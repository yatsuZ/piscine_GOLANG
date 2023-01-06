package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i <= 999; i++ {
		c := (i / 100) % 10
		d := (i / 10) % 10
		u := i % 10
		if u > d && d > c {
			z01.PrintRune(rune(c + 48))
			z01.PrintRune(rune(d + 48))
			z01.PrintRune(rune(u + 48))
			if c == 7 && d == 8 && u == 9 {
				z01.PrintRune('\n')
			} else {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
