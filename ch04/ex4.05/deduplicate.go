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
