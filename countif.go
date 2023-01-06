package piscine

func CountIf(f func(string) bool, tab []string) int {
	resulta := 0
	for i := range tab {
		if f(tab[i]) {
			resulta++
		}
	}
	return resulta
}
