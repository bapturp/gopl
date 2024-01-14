package deduplicate

import (
	"reflect"
	"testing"
)

func TestDeduplicate(t *testing.T) {
	testCases := []struct {
		input   []string
		expects []string
	}{
		{[]string{"alpha", "alpha", "bravo", "charlie", "charlie", "charlie", "delta", "delta"},
			[]string{"alpha", "bravo", "charlie", "delta"}},
		{[]string{""},
			[]string{""}},
	}

	for _, test := range testCases {
		input := make([]string, len(test.input))
		copy(input, test.input)
		result := deduplicate(input)
		if !reflect.DeepEqual(result, test.expects) {
			t.Errorf("Failed deduplicate. input: %v, expects: %v, results: %v", test.input, test.expects, result)
		}
	}
}
