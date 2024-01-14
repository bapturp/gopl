package shadiff

// CountBits counts the number of bits that are different in two SHA256 hashes
func CountBits(h1, h2 [32]byte) int {
	var count int

	for i := 0; i < 32; i++ {
		xorResult := h1[i] ^ h2[i]

		for xorResult != 0 {
			count += int(xorResult & 1)
			xorResult >>= 1
		}
	}
	return count
}
