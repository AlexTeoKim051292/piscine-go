package piscine

func DivMod(a *int, b *int) {

	x := *a % *b
	y := *a / *b

	*a = x
	*b = y
}
