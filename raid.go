package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			if i == 0 || i == x && j == 0 || j == y {
				z01.PrintRune('o')
			} else if i == 0 || i == x {
				z01.PrintRune('-')
			} else if j == 0 || j == y {
				z01.PrintRune('|')
			} else {
				z01.PrintRune(' ')
			}
		}
	}
}
