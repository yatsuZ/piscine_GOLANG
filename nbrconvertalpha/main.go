package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	chifre := 0
	puissance := 1
	if len(os.Args) > 1 {
		nbr := os.Args[1:]
		if nbr[0] == "--upper" {
			nbr = nbr[1:]
			for i := range nbr {
				puissance = 1
				for range nbr[i] {
					puissance = puissance * 10
				}
				chifre = 0
				for j := range nbr[i] {
					puissance = puissance / 10
					chifre = chifre + (int(nbr[i][j])-48)*puissance
				}
				z01.PrintRune(rune(chifre + 64))
			}
		} else if len(nbr) > 0 {
			for I := range nbr {
				puissance = 1
				for range nbr[I] {
					puissance = puissance * 10
				}
				chifre = 0
				for J := range nbr[I] {
					puissance = puissance / 10
					chifre = chifre + (int(nbr[I][J])-48)*puissance
				}
				if chifre >= 27 {
					z01.PrintRune(' ')
				} else {
					z01.PrintRune(rune(chifre + 96))
				}
			}
		}
		z01.PrintRune(10)
	}
}
