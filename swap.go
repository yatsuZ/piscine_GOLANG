package piscine

func Swap(a *int, b *int) {
	inter := *a
	*a = *b
	*b = inter
}
