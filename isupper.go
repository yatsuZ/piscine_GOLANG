package piscine

func IsUpper(s string) bool {
	S := []rune(s)
	for i := range S {
		if S[i] < 64 || S[i] > 91 {
			return false
		}
	}
	return true
}
