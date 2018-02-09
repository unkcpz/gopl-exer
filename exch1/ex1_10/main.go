// Modifiy fetchall to print its output to a file
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[2:] {
		go fetch(url, ch)
	}
	f, err := os.OpenFile(os.Args[1], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	for range os.Args[2:] {
		fmt.Fprintln(f, <-ch)
	}
	fmt.Fprintf(f, "%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("fetchall: %v\n", err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("fetchall: reading %s: %v\n", err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f  %7d  %s", secs, nbytes, url)
}

// After runing `$ ./fetchall log http://gopl.io http://baidu.com` twich
// Getting:
// 2.95     4527  http://gopl.org
// 5.14       81  http://baidu.com
// 5.14s elapsed
// 1.07       81  http://baidu.com
// 1.68     4527  http://gopl.org
// 1.68s elapsed
