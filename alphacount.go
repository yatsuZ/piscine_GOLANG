package piscine

func AlphaCount(s string) int {
	S := []rune(s)
	count := 0
	for i := range s {
		if S[i] >= 'a' && S[i] <= 'z' || S[i] >= 'A' && S[i] <= 'Z' {
			count++
		}
	}
	return count
}
