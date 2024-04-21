package main

import (
	"fmt"
	"strings"
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
