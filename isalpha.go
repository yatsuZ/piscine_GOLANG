package piscine

func IsAlpha(s string) bool {
	S := []rune(s)
	for i := range S {
		if !(S[i] >= '0' && S[i] <= '9' || S[i] >= 'A' && S[i] <= 'Z' || S[i] >= 'a' && S[i] <= 'z') {
			return false
		}
	}
	return true
}
