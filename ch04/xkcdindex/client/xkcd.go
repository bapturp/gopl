package xkcdindex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	xkcdindex "xkcdindex/model"
)

func MakeIndex() {
	var comics []xkcdindex.Comic

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

func GetComic(n int) (xkcdindex.Comic, error) {
	comic := xkcdindex.Comic{}

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
