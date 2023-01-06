package piscine

func SortIntegerTable(table []int) {
	x := 0
	for decompte := range table {
		for i := range table {
			if table[decompte] > table[i] && decompte < i {
				x = table[decompte]
				table[decompte] = table[i]
				table[i] = x
			}
		}
	}
}
