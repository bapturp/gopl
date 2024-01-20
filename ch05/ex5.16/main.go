package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(JoinStrings("|", "foo", "bar", "baz"))
	fmt.Println(JoinStrings("}", "foo"))
	fmt.Println(JoinStrings("|"))

}

func JoinStrings(sep string, vals ...string) string {
	if len(vals) == 0 {
		return ""
	}

	if len(vals) == 1 {
		return vals[0]
	}

	b := bytes.Buffer{}
	b.WriteString(vals[0])
	for _, s := range vals[1:] {
		b.WriteString(sep)
		b.WriteString(s)
	}
	return b.String()
}
