package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Parser struct {
	urls map[string]bool
	mux  sync.Mutex
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {

	p := Parser{urls: make(map[string]bool)}
	p.urls[url] = true
	var crawl func(string, int, Fetcher, chan bool)
	crawl = func(url string, depth int, fetcher Fetcher, c chan bool) {
		if depth <= 0 {
			c <- true
			return
		}

		body, urls, err := fetcher.Fetch(url)

		if err != nil {
			fmt.Println(err)
			c <- true
			return
		}

		fmt.Printf("found: %s %q\n", url, body)

		for _, u := range urls {
			_, exist := p.urls[u]
			if !exist {
				p.mux.Lock()
				p.urls[u] = true
				ch := make(chan bool)
				go crawl(u, depth-1, fetcher, ch)
				p.mux.Unlock()
				<-ch
			}
		}
		c <- true
		return

	}
	c := make(chan bool)
	go crawl(url, depth, fetcher, c)
	<-c
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
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
}
