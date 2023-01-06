package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for m := '0'; m <= '9'; m++ {
		for c := '0'; c <= '9'; c++ {
			for d := '0'; d <= '9'; d++ {
				for u := '0'; u <= '9'; u++ {
					if 10*(m-'0')+c < 10*(d-'0')+u {
						z01.PrintRune(m)
						z01.PrintRune(c)
						z01.PrintRune(' ')
						z01.PrintRune(d)
						z01.PrintRune(u)
						if m == '9' && c == '8' && d == '9' && u == '9' {
							z01.PrintRune('\n')
						} else {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
