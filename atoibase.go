package piscine

func AtoiBase(s string, base string) int {
	nbr := 0
	chifre := 0
	if bonneBase(base) {
		return nbr
	}
	S := []rune(s)
	B := []rune(base)
	puissance := len(S) - 1
	for i := range S {
		for j := range B {
			if S[i] == B[j] {
				chifre = j
			}
		}
		nbr = nbr + (chifre * IterativePower(len(B), puissance))
		puissance--
	}
	return nbr
}
