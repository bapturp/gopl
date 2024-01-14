package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestSha(t *testing.T) {
	tests := []struct {
		input   string
		bitSize int
		want    string
	}{
		{"hello", 256, "f6229a9eff8560fd4bda689a93365c6429770984e1761516adb1c8dc87f1a0d6\n"},
	}
	for _, test := range tests {
		input := strings.NewReader(test.input)
		descr := fmt.Sprintf("sha(%v, %v)", test.bitSize, test.input)
		out = new(bytes.Buffer) // captured output
		if err := sha(&test.bitSize, input); err != nil {
			t.Errorf("%s failed: %v", descr, err)
			continue
		}
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
