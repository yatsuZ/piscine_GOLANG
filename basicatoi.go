package piscine

func BasicAtoi(s string) int {
	nbr := 0
	for i := range s {
		nbr = nbr * 10
		nbr = nbr + int(s[i]-48)
	}
	return nbr
}
