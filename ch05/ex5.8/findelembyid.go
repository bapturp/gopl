package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: findelembyid <url> <id>")
		os.Exit(1)
	}

	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "getting %s: %v\n", url, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing: %v\n", err)
		os.Exit(1)
	}

	id := os.Args[2]
	if found := ElementById(doc, id); found != nil {
		var s []string
		s = append(s, fmt.Sprintf("<%s", found.Data))
		for _, a := range found.Attr {
			s = append(s, fmt.Sprintf("%s='%s'", a.Key, a.Val))
		}
		s = append(s, ">")
		fmt.Printf("Found the id: %s\n", id)
		fmt.Println(strings.Join(s, " "))
	} else {
		fmt.Println("Element NOT found")
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil && pre(n) {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result := forEachNode(c, pre, post); result != nil {
			return result
		}
	}

	if post != nil {
		post(n)
	}

	return nil
}

func ElementById(n *html.Node, id string) *html.Node {
	if n == nil {
		return nil
	}

	// closure to access to the parameter id
	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return false
		}
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
		return false
	}

	return forEachNode(n, pre, nil)
}
