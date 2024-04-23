package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)


func scrapeLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var links []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.HasPrefix(href, "/wiki/") && !strings.Contains(href, "Main_Page")  && !strings.Contains(href, "Special:") && !strings.Contains(href, "Help:") && !strings.Contains(href, "Template:") && !strings.Contains(href, "Category:") && !strings.Contains(href, "Portal:") && !strings.Contains(href, "Wikipedia:"){
			links = append(links, "https://en.wikipedia.org"+href)
		}
	})

	return links, nil
}


func main() {
	url := "https://en.wikipedia.org/wiki/Bandung_Institute_of_Technology"
	links, err := scrapeLinks(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Loop over the links and print them one by one
	for _, link := range links {
		fmt.Println(link) // This prints each link
	}
}