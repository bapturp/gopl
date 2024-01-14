package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	f, err := os.OpenFile("fetchall.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	for range os.Args[1:] {
		if _, err := f.WriteString(<-ch); err != nil {
			fmt.Fprintf(os.Stderr, "writing to file: %s\n", err)
		}
	}

	elapsed := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	if _, err := f.WriteString(elapsed); err != nil {
		fmt.Fprintf(os.Stderr, "writing to file: %s\n", err)
	}
}

// fetch an url
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s\n", secs, nbytes, url)
}
