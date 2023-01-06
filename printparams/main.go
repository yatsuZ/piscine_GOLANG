package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	mots := os.Args
	for i := range mots {
		for j := range mots[i] {
			if mots[i][j] != '/' && mots[i][j] != '.' && i != 0 {
				z01.PrintRune(rune(mots[i][j]))
			}
		}
		if i != 0 {
			z01.PrintRune('\n')
		}
	}
}
