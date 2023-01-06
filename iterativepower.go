package piscine

func IterativePower(nb int, power int) int {
	fix := nb
	if power == 0 {
		return 1
	} else if power < 1 {
		return 0
	}
	for decompte := power; decompte > 1; decompte-- {
		nb = fix * nb
	}
	return nb
}
