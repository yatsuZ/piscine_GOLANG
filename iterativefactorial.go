package piscine

func IterativeFactorial(nb int) int {
	t := 1
	if nb < 0 || nb > 24 {
		return 0
	} else if nb == 0 {
		return 1
	}
	if nb < 25 {
		for i := 1; i <= nb; i++ {
			t = t * i
		}
	}
	return t
}
