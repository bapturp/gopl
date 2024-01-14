package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findelements: %v\n", err)
		os.Exit(1)
	}

	var h = make(HtmlElements)
	visit(h, doc)
	for k, v := range h {
		fmt.Printf("%-10s %d\n", k, v)
	}
}

type HtmlElements map[string]int

func visit(h HtmlElements, n *html.Node) {
	if n != nil {
		if n.Type == html.ElementNode {
			if h[n.Data] == 0 {
				h[n.Data] = 1
			} else {
				h[n.Data]++
			}
		}
		visit(h, n.FirstChild)
		visit(h, n.NextSibling)
	}
}
