package piscine

func ToUpper(s string) string {
	S := []rune(s)
	for i := range s {
		if S[i] >= 97 && S[i] <= 122 {
			S[i] = S[i] - 32
		}
	}
	return string(S)
}
