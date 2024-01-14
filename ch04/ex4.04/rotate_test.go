package rotate

import "testing"

func TestRotate(t *testing.T) {
	testCases := []struct {
		input   []int
		step    int
		expects []int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 1,
			[]int{2, 3, 4, 5, 6, 1}},
		{[]int{1, 2, 3, 4, 5, 6}, 0,
			[]int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3, 4, 5, 6}, -1,
			[]int{1, 2, 3, 4, 5, 6}},
	}

	for _, test := range testCases {
		input := make([]int, len(test.input))
		copy(input, test.input)
		Rotate(input, test.step)
		for i := 0; i < len(test.input); i++ {
			if test.expects[i] != input[i] {
				t.Errorf("Failed rotate. input: %v, step: %v, expects: %v, results: %v", test.input, test.step, test.expects, input)
				break
			}
		}
	}
}
