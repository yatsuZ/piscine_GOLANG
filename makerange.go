package piscine

func MakeRange(min, max int) []int {
	distance := max - min
	if min < max {
		L := make([]int, distance)
		for i := range L {
			L[i] = min + i
		}
		return L
	}
	return nil
}
