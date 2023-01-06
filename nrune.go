package piscine

func NRune(s string, n int) rune {
	S := []rune(s)
	if len(s) < n || n <= 0 {
		return 0
	}
	return S[n-1]
}
