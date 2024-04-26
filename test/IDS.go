package main

import (
	"fmt"
	"time"
)


func DLS(URL string, target string, depth int, visited map[string]bool, path []string, allPaths *[][]string, artikelDiperiksa *int) error {
	
	if depth == 0 {
		*artikelDiperiksa++
		if URL == target {

			*allPaths = append(*allPaths, append([]string(nil), path...))
			return nil
		}

		return nil
	}


	visited[URL] = true
	defer delete(visited, URL) 

	neighbors, err := scrapeLinks(URL)
	if err != nil {
		return err
	}

	for _, neighbor := range neighbors {
		if !visited[neighbor] {
			newPath := append(path, neighbor)
			err := DLS(neighbor, target, depth-1, visited, newPath, allPaths, artikelDiperiksa)
			if err != nil {
				return err
			}
		}
	}

	return nil
}


func IDS(startURL string, targetURL string) ([][]string, int, int, error) {
	artikelDiperiksa := 0
	allPaths := [][]string{} 
	startTime := time.Now()

	for depth := 0; time.Since(startTime).Seconds() <= 300; depth++ {
		visited := make(map[string]bool)
		path := []string{startURL}

		err := DLS(startURL, targetURL, depth, visited, path, &allPaths, &artikelDiperiksa)
		if err != nil {
			return nil, artikelDiperiksa, depth, err
		}

		if len(allPaths) > 0 {
			return allPaths, artikelDiperiksa, depth, nil
		}
	}

	return nil, artikelDiperiksa, 0, fmt.Errorf("Error")
}

