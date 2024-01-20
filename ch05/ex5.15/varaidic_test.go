package variadic

import "testing"

func TestMax(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{5}, 5},
	}

	for _, test := range tests {
		got, _ := Max(test.input...)
		if got != test.want {
			t.Errorf("Max(%q) = %v", test.input, got)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{5}, 5},
	}

	for _, test := range tests {
		got, _ := Min(test.input...)
		if got != test.want {
			t.Errorf("Min(%q) = %v", test.input, got)
		}
	}
}

func TestMax1(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{5}, 5},
	}

	for _, test := range tests {
		got := Max1(test.input[0], test.input[1:]...)
		if got != test.want {
			t.Errorf("Ma1(%q) = %v", test.input, got)
		}
	}
}

func TestMin1(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{5}, 5},
	}

	for _, test := range tests {
		got := Min1(test.input[0], test.input[1:]...)
		if got != test.want {
			t.Errorf("Min1(%q) = %v", test.input, got)
		}
	}
}
