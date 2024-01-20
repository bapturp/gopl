package variadic

import (
	"fmt"
)

func Min(vals ...int) (int, error) {
	if len(vals) < 1 {
		return 0, fmt.Errorf("min require at least one value")
	}

	m := vals[0]
	for _, n := range vals {
		if n < m {
			m = n
		}
	}

	return m, nil
}

func Max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("max require at least one value")
	}
	m := vals[0]
	for _, n := range vals {
		if n > m {
			m = n
		}
	}

	return m, nil
}

func Max1(first int, vals ...int) int {
	m := first
	for _, n := range vals {
		if n > m {
			m = n
		}
	}
	return m
}

func Min1(first int, vals ...int) int {
	m := first
	for _, n := range vals {
		if n < m {
			m = n
		}
	}
	return m
}
