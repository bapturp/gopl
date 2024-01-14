package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("-123456"))
}

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	for i := 0; i < n; i++ {
		if (n-i)%3 == 0 && i != 0 {
			buf.WriteRune(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}
