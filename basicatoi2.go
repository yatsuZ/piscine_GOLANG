package piscine

func BasicAtoi2(s string) int {
	nbr := 0
	for i := range s {
		if rune(s[i]) > 47 && rune(s[i]) < 58 {
			nbr = nbr * 10
			nbr = nbr + int(s[i]-48)
		} else {
			return 0
		}
	}
	return nbr
}
