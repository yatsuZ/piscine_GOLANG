package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	machinbidule := []rune(os.Args[0])
	for c := range machinbidule {
		if machinbidule[c] != '/' && machinbidule[c] != '.' {
			z01.PrintRune(machinbidule[c])
		}
	}
	z01.PrintRune('\n')
}
