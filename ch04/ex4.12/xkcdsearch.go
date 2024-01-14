package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	// MakeIndex()
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

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "xkcdsearch <something>")
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

func MakeIndex() {
	var comics []Comic

	for i := 500; i < 1000; i++ {
		c, err := GetComic(i)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error :%s\n", err)
			continue
		}
		comics = append(comics, c)
	}

	data, err := json.MarshalIndent(comics, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling data: %s\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile("./xkcd-data.json", data, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
		os.Exit(1)
	}
}

func GetComic(n int) (Comic, error) {
	comic := Comic{}

	resp, err := http.Get(fmt.Sprintf("https://xkcd.com/%d/info.0.json", n))
	if err != nil {
		return comic, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return comic, fmt.Errorf("request to get json failed: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return comic, err
	}
	comic.URL = fmt.Sprintf("https://xkcd.com/%d", n)

	return comic, nil
}
