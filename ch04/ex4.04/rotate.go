package rotate

// Rotate rotates the elements in place
func Rotate(s []int, n int) {
	if n < 0 || n >= len(s) {
		return
	}

	temp := make([]int, 0)
	temp = append(temp, s[n:]...)
	temp = append(temp, s[:n]...)
	copy(s, temp)
}
