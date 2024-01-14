package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io"
	"os"
)

var bitSize = flag.Int("s", 256, "Output size (bits), accepted values: 256|384|512")

var in io.Reader = os.Stdin
var out io.Writer = os.Stdout

func main() {
	flag.Parse()
	if err := sha(bitSize, in); err != nil {
		fmt.Fprintf(os.Stderr, "sha: %v\n", err)
		os.Exit(1)
	}
}

func sha(bitSize *int, r io.Reader) error {
	buffer := make([]byte, 1024)
	if _, err := r.Read(buffer); err != nil {
		return fmt.Errorf("reading from stdin: %v\n", err)
	}

	switch *bitSize {
	case 512:
		fmt.Fprintf(out, "%x\n", sha512.Sum512(buffer))
	case 384:
		fmt.Fprintf(out, "%x\n", sha512.Sum384(buffer))
	case 256:
		fmt.Fprintf(out, "%x\n", sha256.Sum256(buffer))
	default:
		return fmt.Errorf("Invalid output size of %d\n", *bitSize)
	}
	return nil
}
