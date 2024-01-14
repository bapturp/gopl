package main

import (
	"testing"
)

func TestComma(t *testing.T) {
	tests := []struct {
		a    string
		want string
	}{
		{"", ""},
		{"1234", "1,234"},
		{"-1234", "-1,234"},
		{"-1234.567", "-1,234.567"},
	}
	for _, tt := range tests { // tt => test table
		t.Run(tt.a, func(t *testing.T) {
			ans := comma(tt.a)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
