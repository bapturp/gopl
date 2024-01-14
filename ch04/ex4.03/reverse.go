package reverse

// Reverse reverses an array of 6 ints in place.
func Reverse(a *[6]int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
