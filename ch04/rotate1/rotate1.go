package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	rotate(s, -1)
	fmt.Println(s)
}

// Rotate the elements forward or backward if n is positive or negative respectively.
func rotate(s []int, n int) {
	if n > 0 {
		n = n % len(s)
		reverse(s)
		reverse(s[n:])
		reverse(s[:n])
	} else if n < 0 {
		n = n * -1 % len(s)
		reverse(s[n:])
		reverse(s[:n])
		reverse(s)
	}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
