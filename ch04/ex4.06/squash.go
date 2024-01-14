// Exercise 4.6:
// Write an in-place function that squashes each run of adjacent Unicode
// spaces (see unicode.IsSpace) in a UTF-8-encoded []byte
// slice into a single ASCII space.

package squash

import "unicode"

func Squash(slice []byte) []byte {
	if len(slice) <= 1 {
		return slice
	}
	var count int // count the removed space

	for i := 0; i < len(slice)-count; i++ {
		if unicode.IsSpace(rune(slice[i])) && unicode.IsSpace(rune(slice[i+1])) {
			copy(slice[i:], slice[i+1:])
			count++
		}
	}

	return slice[:len(slice)-count]
}
