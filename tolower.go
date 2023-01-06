package piscine

func ToLower(s string) string {
	S := []rune(s)
	for i := range s {
		if S[i] >= 65 && S[i] <= 90 {
			S[i] = S[i] + 32
		}
	}
	return string(S)
}
