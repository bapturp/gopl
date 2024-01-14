package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "parsing: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}

var depth int

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	switch n.Type {
	case html.TextNode:
		// indentation is stored as TextNode, but we don't want to render it
		// we check if the content node is only whitespace.
		if !isAllWhitespace(n.Data) {
			fmt.Printf("%*s%s\n", depth*2, "", strings.Trim(n.Data, "\n\t\r "))
		}
	case html.CommentNode:
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", strings.Trim(n.Data, "\n\t\r "))
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		s := n.Data
		if attr := attrToStr(n); attr != "" {
			s += " " + attr
		}
		// there isn't such a thing as self-closing tag like <img /> which should be written <img>
		// https://developer.mozilla.org/en-US/docs/Glossary/Void_element
		fmt.Printf("%*s<%s>\n", depth*2, "", s)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		// If it's a void element, there is no end tag.
		if !isVoidElement(n) {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

// Determine whether a node's text content is entirely whitespace.
// https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model/Whitespace#whitespace_helper_functions
func isAllWhitespace(s string) bool {
	r := regexp.MustCompile(`^[\s]+$`) // whitespace (== [\t\n\f\r ])
	return r.MatchString(s)
}

// determine whether a node is a void element.
func isVoidElement(n *html.Node) bool {
	// A void element is an element in HTML that cannot have any child nodes
	// (i.e., nested elements or text nodes).
	// Void elements only have a start tag; end tags must not be specified for void elements.
	// https://developer.mozilla.org/en-US/docs/Glossary/Void_element

	voidElements := map[string]struct{}{
		"area":   {},
		"base":   {},
		"br":     {},
		"col":    {},
		"embed":  {},
		"hr":     {},
		"img":    {},
		"input":  {},
		"link":   {},
		"meta":   {},
		"param":  {},
		"source": {},
		"track":  {},
		"wbr":    {},
	}

	if _, ok := voidElements[n.Data]; n.Type == html.ElementNode && ok {
		return true
	} else {
		return false
	}
}

// transform the attribute list to a string of space separated key='val'.
func attrToStr(n *html.Node) string {
	var s []string
	if len(n.Attr) != 0 {
		for _, a := range n.Attr {
			s = append(s, fmt.Sprintf("%s='%s'", a.Key, a.Val))
		}
	}
	return strings.Join(s, " ")
}
