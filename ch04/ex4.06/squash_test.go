package squash

import (
	"reflect"
	"testing"
)

func TestSquash(t *testing.T) {
	tests := []struct {
		input    []byte
		expected []byte
	}{
		{
			[]byte{'a', ' ', ' ', 'b'},
			[]byte{'a', ' ', 'b'},
		},
		{
			[]byte{'a', ' ', 'b'},
			[]byte{'a', ' ', 'b'},
		},
		{
			[]byte{' ', ' ', 'b', ' ', ' '},
			[]byte{' ', 'b', ' '},
		},
		{
			[]byte{' ', ' '},
			[]byte{' '},
		},
		{
			[]byte{' '},
			[]byte{' '},
		},
		{
			[]byte{},
			[]byte{},
		},
	}

	for _, test := range tests {
		input := make([]byte, len(test.input))
		copy(input, test.input)
		result := Squash(input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Failed Squash. input: %v, expected: %v, result: %v", test.input, test.expected, result)
		}
	}
}
