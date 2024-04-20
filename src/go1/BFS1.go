package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// Node represents a Wikipedia article node
type Node struct {
	Title string
	URL   string
}

func main() {
	startURL := "https://en.wikipedia.org/wiki/Artificial_intelligence"
	targetURL := "https://en.wikipedia.org/wiki/Power_(physics)"

	startTime := time.Now()

	path, err := BFS(startURL, targetURL)

	elapsed := time.Since(startTime)

	if err != nil {
		log.Fatal(err)
	}

	// print path dengan penomoran
	fmt.Println("\nPath from", startURL, "to", targetURL, "is:")
	for i, node := range path {
		fmt.Println(i+1, ".", node.Title, ":", node.URL)
	}

	fmt.Println("Execution time:", elapsed)
}

func BFS(startURL, targetURL string) ([]Node, error) {
	visited := make(map[string]bool)
	visitedNodes := make(map[string]Node)
	queue := []Node{{"Start", startURL}}

	// Simpan jalur yang ditemukan
	var path []Node

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// // print node
		// fmt.Println("Visiting:", node.Title, ":", node.URL)

		// tandai node sebagai sudah dikunjungi
		visited[node.URL] = true

		// cek apakah node saat ini adalah target
		if node.URL == targetURL {
			// Salin jalur yang ditemukan dari target ke awal
			for node.URL != startURL {
				path = append([]Node{node}, path...)
				node = visitedNodes[node.URL]
			}
			path = append([]Node{{"Start", startURL}}, path...)
			return path, nil
		}

		// dapatkan tetangga-tetangga dari node saat ini
		neighbors, err := getNeighborsFromURL(node.URL)
		if err != nil {
			return nil, err
		}

		// Lakukan pengecekan apakah salah satu tetangga adalah target node
		for _, neighbor := range neighbors {
			if neighbor.URL == targetURL {
				// Jika tetangga adalah target node, tambahkan target node ke jalur terpendek
				path = append(path, node)
				path = append(path, neighbor)
				return path, nil
			}
		}

		// Tambahkan tetangga yang belum dikunjungi ke dalam queue
		for _, neighbor := range neighbors {
			if !visited[neighbor.URL] {
				queue = append(queue, neighbor)
				// Simpan node sebelumnya untuk setiap tetangga
				visitedNodes[neighbor.URL] = node
			}
		}
	}

	return nil, fmt.Errorf("path not found")
}

// scrapping dari wikipedia
func getNeighborsFromURL(URL string) ([]Node, error) {
	// Retrieve konten HTML
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	// Ekstrak URL neighbors
	var neighbors []Node
	var extractURL func(*html.Node)
	extractURL = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" && strings.HasPrefix(attr.Val, "/wiki/") {
					neighborURL := "https://en.wikipedia.org" + attr.Val
					neighborTitle := strings.TrimPrefix(attr.Val, "/wiki/")
					neighbors = append(neighbors, Node{Title: neighborTitle, URL: neighborURL})
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractURL(c)
		}
	}
	extractURL(doc)

	return neighbors, nil
}
