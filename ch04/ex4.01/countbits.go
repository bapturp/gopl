// Exercise 4.1:
// Write a function that counts the number of bits that are different in
// two SHA256 hashes.

package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: countbits <string> <string>")
		os.Exit(1)
	}

	hash1 := sha256.Sum256([]byte(os.Args[1]))
	hash2 := sha256.Sum256([]byte(os.Args[2]))

	bits := CountBits(hash1, hash2)

	fmt.Println(bits)
}

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
