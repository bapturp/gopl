package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("-1234.456"))
	fmt.Println(comma("-1234"))
}

func comma(s string) string {

	if len(s) == 0 {
		return ""
	}

	var integer, fractional string
	if strings.Contains(s, ".") {
		split := strings.Split(s, ".")
		integer = split[0]
		fractional = split[1]
	} else {
		integer = s
	}

	var buf bytes.Buffer

	// write minus sign
	if integer[0] == '-' {
		buf.WriteRune('-')
		integer = integer[1:]
	}

	// write number with thousand separator
	thousandSep := ','
	n := len(integer)
	for i := 0; i < n; i++ {
		if (n-i)%3 == 0 && i != 0 {
			buf.WriteRune(thousandSep)
		}

		buf.WriteByte(integer[i])
	}

	decimalSep := '.'
	if fractional != "" {
		buf.WriteRune(decimalSep)
		for i := range fractional {
			buf.WriteByte(fractional[i])
		}
	}

	return buf.String()
}
