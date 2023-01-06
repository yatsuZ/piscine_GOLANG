package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	nbr := str2nbr(os.Args[1])
	tab(nbr)
}

func tab(nbr int) {
	for i := 1; i < 10; i++ {
		atoibeta(i)
		z01.PrintRune('*')
		atoibeta(nbr)
		z01.PrintRune('=')
		atoibeta(i * nbr)
		z01.PrintRune(10)
	}
}

func atoibeta(nbr int) {
	s := ""
	for nbr != 0 {
		s = string(rune((nbr%10)+48)) + s
		nbr = nbr / 10
	}
	for _, j := range s {
		z01.PrintRune(rune(j))
	}
}

func str2nbr(s string) int {
	S := []rune(s)
	if len(S) == 1 {
		return int(rune(S[0]) - 48)
	}
	nbr := 0
	for i := range S {
		nbr *= 10
		nbr += int(rune(S[i] - 48))
	}
	return nbr
}
