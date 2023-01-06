package piscine

func IsPrintable(s string) bool {
	S := []rune(s)
	for i := range S {
		if S[i] < 32 || s[i] == 127 {
			return false
		}
	}
	return true
}
