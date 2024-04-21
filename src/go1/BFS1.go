package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Node untuk artikel Wikipedia
type Node struct {
	Title    string
	URL      string
	Parent   *Node
	Children []*Node
}

// BFS
func BFS(startURL, targetURL string) ([]*Node, int, int, error) {
	visited := make(map[string]bool)
	var pathToTarget []*Node
	var articlesChecked, articlesTraversed int

	startTitle := strings.TrimPrefix(startURL, "https://en.wikipedia.org/wiki/")

	queue := []*Node{{Title: startTitle, URL: startURL}}

	articlesChecked++

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if visited[node.URL] {
			continue
		}

		visited[node.URL] = true

		articlesTraversed++

		if node.URL == targetURL {
			pathToTarget = buildPathToTarget(node)
			return pathToTarget, articlesChecked, articlesTraversed, nil
		}

		neighbors, err := getNeighborsFromURL(node.URL)
		if err != nil {
			return nil, articlesChecked, articlesTraversed, err
		}

		for _, neighbor := range neighbors {
			if neighbor.URL == targetURL {
				return buildPathToTarget(&Node{Title: neighbor.Title, URL: targetURL, Parent: node}), articlesChecked, articlesTraversed, nil
			}

			neighbor.Parent = node
			node.Children = append(node.Children, neighbor)
			if !visited[neighbor.URL] {
				queue = append(queue, neighbor)
				articlesChecked++
			}
		}
	}

	return nil, articlesChecked, articlesTraversed, fmt.Errorf("path not found")
}

// Membuat rute penjelajahan
func buildPathToTarget(target *Node) []*Node {
	var path []*Node
	current := target
	for current != nil {
		path = append([]*Node{current}, path...)
		current = current.Parent
	}
	return path
}

func main() {
	startURL := "https://en.wikipedia.org/wiki/French-suited_playing_cards"
	targetURL := "https://en.wikipedia.org/wiki/Indian_Premier_League"

	// startURL := "https://en.wikipedia.org/wiki/Physics"
	// targetURL := "https://en.wikipedia.org/wiki/Indian_Premier_League"

	startTime := time.Now()

	path, articlesChecked, articlesTraversed, err := BFS(startURL, targetURL)

	elapsed := time.Since(startTime)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nPath dari", startURL, "ke", targetURL, "adalah:")
	for i, node := range path {
		fmt.Println(i+1, ".", node.Title, ":", node.URL)
	}

	fmt.Println("\nArtikel diperiksa:", articlesChecked)
	fmt.Println("Artikel dikunjungi:", articlesTraversed)
	fmt.Println("Waktu eksekusi:", elapsed)
}
