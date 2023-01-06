package piscine

func TrimAtoi(s string) int {
	nbr := 0
	boulé := true
	CMORT := 0
	for i := range s {
		if rune(s[i]) >= '0' && rune(s[i]) < 58 || rune(s[i]) == '-' && CMORT == 0 {
			if rune(s[i]) == '-' {
				boulé = false
			} else if rune(s[i]) == '+' {
				boulé = true
			} else {
				CMORT = CMORT + 1
				nbr = nbr * 10
				nbr = nbr + int(s[i]-48)
			}
		}
	}
	if boulé == false {
		return (-1 * nbr)
	} else {
		return nbr
	}
}
