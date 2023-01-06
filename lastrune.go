package piscine

func LastRune(s string) rune {
	X := []rune(s)
	return X[len(X)-1]
}
