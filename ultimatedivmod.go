package piscine

func UltimateDivMod(a *int, b *int) {

	x := *a % *b
	y := *a / *b

	*a = y
	*b = x
}
