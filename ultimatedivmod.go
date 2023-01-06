package piscine

func UltimateDivMod(a *int, b *int) {
	intermidiere := *a
	*a = *a / *b
	*b = intermidiere % *b
}
