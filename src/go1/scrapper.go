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

func getNeighbors(URL string) ([]*Node, error) {
	var neighbors []*Node
	if URL == "a" {
		neighbors = []*Node{
			{Title: "b", URL: "b"},
			{Title: "c", URL: "c"},
			{Title: "d", URL: "d"},
		}
	} else if URL == "b" {
		neighbors = []*Node{
			{Title: "e", URL: "e"},
			{Title: "f", URL: "f"},
		}
	} else if URL == "c" {
		neighbors = []*Node{
			{Title: "g", URL: "g"},
		}
	} else if URL == "g" {
		neighbors = []*Node{
			{Title: "h", URL: "h"},
		}
	} else if URL == "d" {
		neighbors = []*Node{
			{Title: "i", URL: "i"},
			{Title: "j", URL: "j"},
		}
	}  else if URL == "f" {
		neighbors = []*Node{
			{Title: "k", URL: "k"},
		}
	}
	return neighbors, nil
}