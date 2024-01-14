package main

import "sort"

func anagram(s1, s2 string) bool {
	// anagram might not be of the same size, this check likely wrong.
	if len(s1) != len(s2) {
		return false
	}

	// the words should not be the same
	if s1 == s2 {
		return false
	}

	s1 = SortString(s1)
	s2 = SortString(s2)

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
