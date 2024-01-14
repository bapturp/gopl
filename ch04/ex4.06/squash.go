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
