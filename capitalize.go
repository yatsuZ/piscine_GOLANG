package piscine

func Capitalize(s string) string {
	maj := true
	S := []rune(s)
	for i := range S {
		if S[i] >= '0' && S[i] <= '9' || S[i] >= 'A' && S[i] <= 'Z' || S[i] >= 'a' && S[i] <= 'z' {
			if S[i] >= 'a' && S[i] <= 'z' && maj == true {
				S[i] = S[i] - 32
			} else if S[i] >= 'A' && S[i] <= 'Z' && maj == false {
				S[i] = S[i] + 32
			}
			maj = false
		} else {
			maj = true
		}
	}
	return string(S)
}
