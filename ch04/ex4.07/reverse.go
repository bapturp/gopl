// Exercise 4.7:
// Modify reverse to reverse the characters of a
// []byte slice that represents a UTF-8-encoded string, in place.
// Can you do it without allocating new memory?

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	text := []byte{0xf0, 0x9f, 0x98, 0x80, 0xf0, 0x9f, 0x98, 0x8E}
	fmt.Println(utf8.RuneCount(text))
	fmt.Println(utf8.Valid(text[:4]))
}

// Reverse reverses an []byte slice of utf-8 characters
func Reverse(slice []byte) []byte {
	n := utf8.RuneCount(slice)
	if n <= 1 {
		return slice
	}

	fmt.Println(n)
	i := utf8.RuneCount(slice) / 2 // the number of permutation

	for j := 0; j < i; j++ {

	}

	return []byte{}
}
