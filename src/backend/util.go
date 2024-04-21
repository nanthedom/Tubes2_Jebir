package backend

import (
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
		if exists && strings.HasPrefix(href, "/wiki/") {
			links = append(links, "https://en.wikipedia.org"+href)
		}
	})

	return links, nil
}

func buildPathToTarget(target *Node) []string {
	var path []string
	current := target
	for current != nil {
		path = append([]string{current.Url}, path...)
		current = current.Prev
	}
	return path
}
