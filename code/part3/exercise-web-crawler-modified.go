package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Add
func CrawlRecursive(url string, depth int, fetcher Fetcher, quit chan bool, visitedUrls map[string]bool) {
	if depth <= 0 {
		quit <- true
		return
	}

	didIt, hasIt := visitedUrls[url]

	if didIt && hasIt {
		quit <- true
		return

	} else {
		visitedUrls[url] = true
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		quit <- true
		return
	}
	fmt.Printf("found URL: %s ; title: %q\n", url, body)

	childrenQuit := make(chan bool)
	for _, childrenUrl := range urls {
		go CrawlRecursive(childrenUrl, depth-1, fetcher, childrenQuit, visitedUrls)
		<-childrenQuit
	}

	quit <- true
}

func Crawl(url string, depth int, fetcher Fetcher) {
	quit := make(chan bool)
	visitedUrls := map[string]bool{url: false}

	go CrawlRecursive(url, depth, fetcher, quit, visitedUrls)

	<-quit
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
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
