package main

import (
	"os"

	"github.com/01-edu/z01"
)

func even(nbr int) int {
	for i := nbr; i >= 2; i = i - 2 {
		nbr = nbr - 2
	}
	return nbr
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if even(nbr) == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	OddMsg := "I have an odd number of arguments"
	EvenMsg := "I have an even number of arguments"
	lengthOfArg := len(os.Args[1:]) - 1
	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
