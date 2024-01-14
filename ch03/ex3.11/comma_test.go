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

	for _, test := range tests { // tt => test table
		t.Run(test.a, func(t *testing.T) {
			ans := comma(test.a)
			if ans != test.want {
				t.Errorf("got %s, want %s", ans, test.want)
			}
		})
	}
}
