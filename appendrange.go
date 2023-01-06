package piscine

func AppendRange(min, max int) []int {
	var l []int
	for i := min; i < max; i++ {
		l = append(l, i)
	}
	return l[0:]
}
