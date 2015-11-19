package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan int) {

	// TODO: Fetch URLs in parallel. --> DONE
	// TODO: Don't fetch the same URL twice. --> DONE
	// This implementation doesn't do either:
	if depth <= 0 {
		//we're done crawling in this "chain"
		//decrement counter
		ch <- (-1)
		return
	}

	//has the page been crawled?
	if _, ok := crawled[url]; ok == true {
		//decrement counter
		ch <- (-1)
		return
	}

	body, urls, err := fetcher.Fetch(url)

	//note we've crawled page
	crawled[url] = true

	//fmt.Println("Go:", url, len(urls))
	ch <- (len(urls) - 1)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, ch)
	}
	return
}

var crawled map[string]bool = make(map[string]bool)

func main() {

	ch := make(chan int)
	n := 0

	go Crawl("http://golang.org/", 4, fetcher, ch)

	for {

		n += <- ch

		if n <= 0 {
			//there is left more to crawl
			close(ch)
			break
		}
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
			"http://google.com/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://google.com/": &fakeResult{
		"Search",
		[]string{

		},
	},
}
