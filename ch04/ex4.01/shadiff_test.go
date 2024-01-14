package shadiff

import (
	"crypto/sha256"
	"testing"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		input []string
		want  int
	}{
		{[]string{"foo", "bar"}, 150},
		{[]string{"foo", "foo"}, 0},
	}

	for _, test := range tests {
		h1 := sha256.Sum256([]byte(test.input[0]))
		h2 := sha256.Sum256([]byte(test.input[1]))
		if got := CountBits(h1, h2); got != test.want {
			t.Errorf("CountBits(%v) = %v", test.input, test.want)
		}
	}
}
