package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getNeighborsFromURL(URL string) ([]*Node, error) {
	res, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("failed to retrieve URL: %s", res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	neighbors := []*Node{}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && strings.HasPrefix(href, "/wiki/") {
			neighborURL := "https://en.wikipedia.org" + href
			neighborTitle := strings.TrimPrefix(href, "/wiki/")
			neighbors = append(neighbors, &Node{Title: neighborTitle, URL: neighborURL})
		}
	})

	return neighbors, nil
}
