package main

import (
	"fmt"
)

func k(list []string) string                    { return fmt.Sprintf("%q", list) }
func Add(list []string, m map[string]int)       { m[k(list)]++ }
func Count(list []string, m map[string]int) int { return m[k(list)] }

func main() {
	s1 := []string{"alpha", "bravo", "charlie"}
	s2 := []string{"delta", "echo", "foxtrot"}
	var m = make(map[string]int)
	// fmt.Printf("%q\n", s1)

	Add(s1, m)
	Add(s1, m)
	Add(s2, m)

	fmt.Println(Count(s1, m))
	fmt.Println(Count(s2, m))
	fmt.Println(len(m))

}
