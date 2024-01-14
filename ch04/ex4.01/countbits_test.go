package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		args     []string
		expected string
	}{
		{[]string{"foo", "bar"}, "150\n"},
		{[]string{"foo", "foo"}, "0\n"},
		{[]string{"foo", "bar", "baz"}, "usage: countbits <string> <string>\nexit status 1\n"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Args: %v", test.args), func(t *testing.T) {
			output, _ := runMain(test.args...)
			// if err != nil {
			// 	t.Fatalf("Error running main: %v", err)
			// }
			if output != test.expected {
				t.Errorf("Error testcount. args: %v, expected: %v, result: %v", test.args, test.expected, output)
			}
		})
	}
}

func runMain(args ...string) (string, error) {
	cmd := exec.Command("go", "run", "countbits.go")
	cmd.Args = append(cmd.Args, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}
