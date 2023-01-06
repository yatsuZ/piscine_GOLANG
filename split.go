package piscine

func Split(s, sep string) []string {
	mots := ""
	najoutepaspendant := 0
	var L []string
	t := len(s) - len(sep)
	for i := range s {
		if sep != string(s[t]) && len(s)-len(sep) < i {
			mots = mots + string(s[i])
		}
		if najoutepaspendant <= 0 && len(s)-len(sep) >= i {
			if sep != s[i:len(sep)+i] {
				mots = mots + string(s[i])
			} else if sep == s[i:len(sep)+i] {
				najoutepaspendant = len(sep)
				L = append(L, mots)
				mots = ""
			}
		}
		najoutepaspendant--
	}
	L = append(L, mots)
	return L
}
