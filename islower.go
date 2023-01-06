package piscine

func IsLower(s string) bool {
	S := []rune(s)
	for i := range S {
		if S[i] < 97 || S[i] > 122 {
			return false
		}
	}
	return true
}
