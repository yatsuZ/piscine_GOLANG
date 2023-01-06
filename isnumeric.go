package piscine

func IsNumeric(s string) bool {
	S := []rune(s)
	for i := range S {
		if !(S[i] >= '0' && S[i] <= '9') {
			return false
		}
	}
	return true
}
