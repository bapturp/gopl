package reverseutf8

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		input    []byte
		expected []byte
	}{
		{
			[]byte{0xf0, 0x9f, 0x98, 0x80, 0xf0, 0x9f, 0x98, 0x8E}, // 😀😎
			[]byte{0xf0, 0x9f, 0x98, 0x8E, 0xf0, 0x9f, 0x98, 0x80}, // 😎😀
		},
	}

	for _, test := range tests {
		input := make([]byte, len(test.input))
		copy(input, test.input)
		result := Reverse(input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Failed Reverse. input: %v, expected: %v, result: %v", test.input, test.expected, result)
		}
	}
}
