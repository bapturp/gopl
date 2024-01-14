package main

import (
	"testing"
)

func TestAnagram(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want bool
	}{
		{"earth", "heart", true},
		{"poney", "poney", false},
		{"debit card", "bad credit", true},
	}
	for _, tt := range tests { // tt => test table
		testName := tt.a
		t.Run(testName, func(t *testing.T) {
			ans := anagram(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
