package piscine

func Map(f func(int) bool, a []int) []bool {
	L := make([]bool, len(a))
	for i := range L {
		L[i] = f(a[i])
	}
	return L
}
