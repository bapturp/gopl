package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	zero(&c1)
}

// We use a pointer to modify the original array since arrays are passed by copy and not reference.
// Using a pointer allow us to midify the original array as intended.
func zero(p *[32]byte) {
	for i := range p {
		p[i] = 0
	}
}
