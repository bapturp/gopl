package main

import (
	"reflect"
	"testing"
)

func TestJoinStrings(t *testing.T) {
	tests := []struct {
		vals []string
		sep  string
		want string
	}{
		{[]string{""}, " ", ""},
		{[]string{"foo"}, " ", "foo"},
		{[]string{"foo", "bar"}, " ", "foo bar"},
	}

	for _, test := range tests {
		got := JoinStrings(test.sep, test.vals...)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("JoinStrings(%q, %q) = %s\n", test.sep, test.vals, got)
		}
	}
}
