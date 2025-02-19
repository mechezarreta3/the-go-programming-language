// Fetchall exercise 1-10
// fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a go goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	filename := ""

	if trimmed_url, ok := strings.CutPrefix(url, "https://"); ok {
		filename = trimmed_url + time.Since(start).String() + ".txt"
	} else if trimmed_url, ok := strings.CutPrefix(url, "http://"); ok {
		filename = trimmed_url + ".txt"
	} else {
		filename = url + ".txt"
	}

	file, err := os.Create(filename)
	if err != nil {
		ch <- fmt.Sprintf("error creating file %s: %v", filename, err)
		return
	}
	defer file.Close()
	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
