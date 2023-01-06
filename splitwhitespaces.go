package piscine

func SplitWhiteSpaces(s string) []string {
	mot := ""
	espace := true
	var L []string
	for _, i := range s {
		if !(rune(i) == ' ' || rune(i) == '\n' || rune(i) == '	') {
			mot = mot + string(i)
			espace = false
		} else {
			if espace == false && mot != "" {
				L = append(L, mot)
				mot = ""
			} else {
				espace = true
			}
		}
	}
	if espace == false && mot != "" {
		L = append(L, mot)
	}
	return L
}
