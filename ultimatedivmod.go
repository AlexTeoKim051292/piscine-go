package piscine

func DivMod(a *int, b *int) {
	
	b: = a % b	
	a: = a / b

	*a = a
	*b = b
}
