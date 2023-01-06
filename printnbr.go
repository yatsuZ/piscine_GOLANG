package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	p := 1
	T := 0
	S := ""
	if n < 0 {
		S = S + "-"
		for Z := n; Z < -10; Z = Z / 10 {
			p = p * 10
		}
		for v := p; v >= 1; v = v / 10 {
			T = (n / p)
			n = n - (T * p)
			p = p / 10
			S = S + string((rune((T * -1) + 48)))
		}
	} else {
		for Z := n; Z > 10; Z = Z / 10 {
			p = p * 10
		}
		for v := p; v >= 1; v = v / 10 {
			T = (n / p)
			n = n - (T * p)
			p = p / 10
			S = S + string((rune(T + 48)))
		}
	}
	for _, i := range S {
		z01.PrintRune(rune(i))
	}
}
