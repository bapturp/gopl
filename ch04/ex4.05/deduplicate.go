// Exercise 4.5:
// Write an in-place function to eliminate adjacent duplicates
// in a []string slice.

package deduplicate

// dedupplicate duplicates adjacent strings
func deduplicate(slice []string) []string {
	if len(slice) == 0 {
		return slice
	}

	var i int // index of the last written string
	for _, s := range slice {
		if s != slice[i] {
			i++
			slice[i] = s
		}
	}

	return slice[:i+1]
}
