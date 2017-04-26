package main

import (
	"fmt"
	"sync"
)

type ContentCache struct {
	c   map[string]string // key: url, value: body
	mux sync.Mutex
}

func (c *ContentCache) Get(url string) (body string, ok bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	body, ok = c.c[url]
	return
}

func (c *ContentCache) Put(url string, body string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.c[url] = body
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *ContentCache) {
	fmt.Printf("Start crawling. depth=%d, url=%s\n", depth, url)

	if depth <= 0 {
		return
	}

	// Don't fetch the same URL twice.
	if _, ok := cache.Get(url); ok {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	// Add url to cache
	cache.Put(url, body)

	// Try to fetch liked URLs
	done := make(chan int)
	for _, u := range urls {
		go func (url string) {
			Crawl(url, depth - 1, fetcher, cache)
			done <- 0
		}(u)
	}
	for range urls {
		<-done
	}
	return
}

func main() {
	Crawl("http://golang.org/", 4, fetcher, &ContentCache{c: make(map[string]string)})
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
