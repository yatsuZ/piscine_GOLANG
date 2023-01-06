package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func PrintNbr(n int) {
	p := 1
	T := 0
	if n < 0 {
		z01.PrintRune('-')
		for Z := n; Z < -10; Z = Z / 10 {
			p = p * 10
		}
		for v := p; v >= 1; v = v / 10 {
			T = (n / p)
			n = n - (T * p)
			p = p / 10
			z01.PrintRune(rune((T * -1) + 48))
		}
	} else {
		for Z := n; Z > 10; Z = Z / 10 {
			p = p * 10
		}
		for v := p; v >= 1; v = v / 10 {
			T = (n / p)
			n = n - (T * p)
			p = p / 10
			z01.PrintRune(rune(T + 48))
		}
	}
}

func main() {
	points := &point{}

	setPoint(points)

	affichageX := "x = "
	affichageY := " y = "
	for _, i := range affichageX {
		z01.PrintRune(rune(i))
	}
	PrintNbr(points.x)
	z01.PrintRune(',')
	for _, j := range affichageY {
		z01.PrintRune(j)
	}
	PrintNbr(points.y)
	z01.PrintRune('\n')
}
