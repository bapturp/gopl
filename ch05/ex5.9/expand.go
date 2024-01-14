// Exercise 5.9:
// Write a function expand(s string, f func(string) string) string that
// replaces each substring “$foo” within s by the text returned
// by f("foo").

package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "There's starman waiting in the $foo."
	r := replacer("foo", "sky")

	fmt.Println(expand(text, r))
}

// returns a function that replaces the string `s` by `t`
func replacer(s string, t string) func(string) string {
	return func(input string) string {
		r := regexp.MustCompile(`\$` + s)
		return r.ReplaceAllString(input, t)
	}
}

// applies the given replacement function `f` to the string `s`
func expand(s string, f func(string) string) string {
	return f(s)
}
