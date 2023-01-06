package piscine

func Join(strs []string, sep string) string {
	if len(strs) <= 1 {
		return string(strs[0])
	} else {
		return Join(strs[:len(strs)/2], sep) + sep + Join(strs[len(strs)/2:], sep)
	}
}
