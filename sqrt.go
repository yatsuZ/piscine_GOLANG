package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	} else if nb > 1 {
		for n := 0; n <= nb/n; n++ {
			if n*n == nb {
				return n
			}
		}
	}
	return 0
}
