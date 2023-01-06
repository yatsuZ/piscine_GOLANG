package piscine

func Premier(nbr int) []int {
	nombre := 0
	var a []int
	for i := 2; i <= nbr; i++ {
		for j := 2; j <= nbr; j++ {
			if i%j == 0 && i != j {
				nombre = 1
			}
		}
		if nombre == 0 {
			a = append(a, i)
		}
		nombre = 0
	}
	return a
}

func PremierDiv(nbr int) []int {
	p := nbr
	a := Premier(nbr)
	var div []int
	for _, i := range a {
		for j := 1; j <= p; j++ {
			if p%i == 0 {
				div = append(div, i)
				p = p / i
			}
		}
	}
	return div
}
