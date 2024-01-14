package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	visit(doc)

}

func visit(n *html.Node) {
	if n != nil {
		if n.Type == html.TextNode {
			s := strings.TrimSpace(n.Data)
			if s != "" {
				fmt.Printf("%s\n", s)
			}
		}
		if n.Data != "script" && n.Data != "style" {
			visit(n.FirstChild)
			visit(n.NextSibling)
		}
	}
}
