package main

//Fetch prints the content found at a URL
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //start a go routine
	}
	for range os.Args[1:] {
		data := (<-ch)
		f.WriteString(data)
	}
	fmt.Printf("%.2fs elapsed \n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	b, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close() //don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2fs\n %d\n %s\n", secs, b, url)
}

/*
Goroutine is a concurrent function execution. Channel is a communication mechanism
that allows one goroutine to pass values of a specified type to another goroutine
	-> main runs in a go routine and the go statement makes another go routine

main function creates a channel of strings using make.
	-> fetch is called async in the new go routine

When one goroutine attempts a send or receive on a channel, it blocks until another goroutine
attempts the corresponding send or receive operation at which point the
value is transferred and both go rountines proceed.

Each Fetch sends a value (ch <- expression) on the channel ch, and main receives all of them.

To answer the question, when I run the command again and a second time, it does about .05 to .30 seconds
better than when I ran it originally.

sumerjoshi@sumers-mbp main % go run chp1/fetchingURLSConcurrently/fetchAll.go http://google.com http://twitter.com http://facebook.com
0.256002s   20471 http://google.com
0.496464s    2610 http://twitter.com
0.626132s   68174 http://facebook.com
0.63s elapsed
sumerjoshi@sumers-mbp main % go run chp1/fetchingURLSConcurrently/fetchAll.go http://google.com http://twitter.com http://facebook.com
0.191304s   20493 http://google.com
0.419031s    2610 http://twitter.com
0.537965s   68171 http://facebook.com
0.54s elapsed

In total, we see it did ~.10 seconds better than the first run.



*/
