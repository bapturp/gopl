package main

import (
	"slices"
)

func anagram(s1, s2 string) bool {
	// the words should not be the same
	if s1 == s2 {
		return false
	}

	ss1 := []rune(s1)
	ss2 := []rune(s2)

	slices.Sort([]rune(ss1))
	slices.Sort([]rune(ss2))

	return string(ss1) == string(ss2)
}
