package main

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"
)

type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func main() {
	var s Sitemapindex
	var n News
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)

	xml.Unmarshal(bytes, &s)

	for _, Location := range s.Locations {
		link := strings.Split(Location, "/")
		to := link[len(link)-1]

		resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/" + to)
		bytes, _ := ioutil.ReadAll(resp.Body)

		xml.Unmarshal(bytes, &n)
	}
}
