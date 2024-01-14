// Exercise 4.2:
// Write a program that prints the SHA256 hash of its standard input
// by default but supports a command-line flag to print the
// SHA384 or SHA512 hash instead.

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	bitSize := flag.Int("s", 256, "Output size (bits), accepted values: 256|384|512")
	flag.Parse()

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "reading from stdin: %v\n", err)
		os.Exit(1)
	}

	switch *bitSize {
	case 512:
		fmt.Printf("%x\n", sha512.Sum512(data))
	case 384:
		fmt.Printf("%x\n", sha512.Sum384(data))
	case 256:
		fmt.Printf("%x\n", sha256.Sum256(data))
	default:
		fmt.Fprintf(os.Stderr, "Invalid output size of %d\n", *bitSize)
		os.Exit(1)
	}
}
