package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines("stdin", os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(arg, f, counts)
			f.Close()
		}
	}
	for file, dup := range counts {
		for line, n := range dup {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", file, n, line)
			}
		}

	}
}

func countLines(filename string, f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if _, ok := counts[filename]; !ok {
			counts[filename] = make(map[string]int)
		}
		counts[filename][input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
