package main

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

// checkExcludedKeywords checks if any excluded keyword is in the given string.
func checkExcludedKeywords(str string, keywords []string) bool {
	for _, keyword := range keywords {
		if strings.Contains(str, keyword) {
			return true // If it contains any of the excluded keywords, return true
		}
	}
	return false // Otherwise, return false
}

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
	excludedKeywords := []string{"Main_Page", "Special:", "Help:", "Template:", "Category:", "Portal:", "Wikipedia:"}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.HasPrefix(href, "/wiki/") && !checkExcludedKeywords(href, excludedKeywords) {
			links = append(links, "https://en.wikipedia.org"+href)
		}
	})

	return links, nil
}

func main() {
	url := "https://en.wikipedia.org/wiki/Main_Page"
	links, err := scrapeLinks(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, link := range links {
		fmt.Println(link)
	}
}