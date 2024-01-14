package reverse

import "testing"

func TestReverse(t *testing.T) {
	testCases := []struct {
		input [6]int
		want  [6]int
	}{
		{
			[6]int{1, 2, 3, 4, 5, 6},
			[6]int{6, 5, 4, 3, 2, 1},
		},
		{
			[6]int{-3, -2, -1, 0, 1, 2},
			[6]int{2, 1, 0, -1, -2, -3},
		},
	}

	for _, test := range testCases {
		got := test.input
		Reverse(&got)
		for i := 0; i < len(test.input); i++ {
			if test.want[i] != got[i] {
				t.Errorf("Failed Reverse. input: %v, want: %v, result: %v", test.input, test.want, got)
				break
			}
		}
	}
}
