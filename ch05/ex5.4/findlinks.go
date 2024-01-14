// Exercise 5.4:
// Extend the visit function so that it extracts other kinds of
// links from the document, such as images, scripts, and style sheets.”

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n != nil {
		if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "img" || n.Data == "script" || n.Data == "link") {
			for _, a := range n.Attr {
				if a.Key == "src" || a.Key == "href" || a.Key == "rel" {
					links = append(links, a.Val)
				}
			}
		}
		links = visit(links, n.FirstChild)
		links = visit(links, n.NextSibling)
	} else {
		return links
	}

	return links
}
