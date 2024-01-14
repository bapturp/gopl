package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
	URL        string
}

func main() {
	file, err := os.ReadFile("./xkcd-data.json")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading the file: %s\n", err)
		os.Exit(1)
	}

	var comics []Comic
	if err := json.Unmarshal(file, &comics); err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshaling the file: %s\n", err)
		os.Exit(1)
	}

	if len(os.Args) < 1 {
		fmt.Fprintln(os.Stderr, "Usage xkcdsearch <term>")
		os.Exit(1)
	}

	term := os.Args[1]

	for _, c := range comics {
		if strings.Contains(strings.ToLower(c.Title), strings.ToLower(term)) {
			fmt.Printf("XKCD %d\n", c.Num)
			fmt.Printf("\t%s\n", c.Title)
			fmt.Printf("\t%s\n", c.URL)
		}
	}
}
