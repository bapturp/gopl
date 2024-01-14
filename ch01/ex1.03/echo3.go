package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(EchoConcat(os.Args[1:], " "))
}

func EchoConcat(args []string, sep string) string {
	var s string
	for _, arg := range args {
		s += sep + arg
	}
	return s
}

func EchoJoin(args []string, sep string) string {
	return strings.Join(args, sep)
}
