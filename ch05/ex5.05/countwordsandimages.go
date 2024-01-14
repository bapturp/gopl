package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "countwordsandimages: %v\n", err)
			continue
		}
		fmt.Printf("Site: %s\nWords: %d\nImages: %d\n", url, words, images)
	}
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return 0, 0, err
	}

	words, images = countWordsAndImages(0, 0, doc)

	return words, images, nil
}

func countWordsAndImages(words, images int, n *html.Node) (int, int) {
	if n != nil {
		if n.Type == html.TextNode {
			s := strings.TrimSpace(n.Data)
			words += len(strings.Split(s, " "))
		} else if n.Type == html.ElementNode && n.Data == "img" {
			images++
		}
		words, images = countWordsAndImages(words, images, n.FirstChild)
		words, images = countWordsAndImages(words, images, n.NextSibling)
	} else {
		return words, images
	}
	return words, images
}
