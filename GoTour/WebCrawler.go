package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	visited := visitedUrls{urls: make(map[string]int)}
	go CrawlRec(url, depth, fetcher, visited)
	return
}

func CrawlRec(url string, depth int, fetcher Fetcher, visited visitedUrls) {
	visited.mux.Lock()
	if _, ok := visited.urls[url]; depth <= 0 || ok {
		visited.mux.Unlock()
		return
	}

	// adding url to visited map
	visited.urls[url] = 2137
	visited.mux.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go CrawlRec(u, depth-1, fetcher, visited)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
	time.Sleep(time.Second)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

type visitedUrls struct {
	urls map[string]int
	mux  sync.Mutex
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

/* correct answer (with goroutines might look different)
found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
not found: https://golang.org/cmd/
found: https://golang.org/pkg/fmt/ "Package fmt"
found: https://golang.org/pkg/os/ "Package os"
*/
