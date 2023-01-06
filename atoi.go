package piscine

func Atoi(s string) int {
	nbr := 0
	boulé := true
	for i := range s {
		if rune(s[i]) > 47 && rune(s[i]) < 58 || rune(s[0]) == '-' && i == 0 || rune(s[0]) == '+' && i == 1 {
			if rune(s[i]) == '-' {
				boulé = false
			} else if rune(s[i]) == '+' {
				boulé = true
			} else {
				nbr = nbr * 10
				nbr = nbr + int(s[i]-48)
			}
		} else if !(rune(s[i]) == ' ') {
			return 0
		}
	}
	if boulé == false {
		return (-1 * nbr)
	} else {
		return nbr
	}
}
