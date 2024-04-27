package backend

import (
	"fmt"
	"strings"
	"time"
)

func DLSAllPath(URL string, target string, depth int, visited map[string]bool, path []string, visitedPath map[string]bool, allPaths *[][]string, artikelDiperiksa *int) error {
	if depth == 0  && URL == target{ //menambahlan kedalam path jika url == target
		*artikelDiperiksa++
		newPath := append([]string(nil), path...)
		pathString := strings.Join(newPath, "->")
		if !visitedPath[pathString] {
			visitedPath[pathString] = true
			*allPaths = append(*allPaths, newPath)
			return nil
		}
	}

	if depth == 0 { //menambahkan artikel diperiksa saat kedalamn 0 (artikel yang baru)
		*artikelDiperiksa++
	}

	if depth > 0 {
		visited[URL] = true
		defer delete(visited, URL)

		neighbors, err := scrapeLinks(URL)
		if err != nil {
			return err
		}

		for _, neighbor := range neighbors { //menerapkan DLS untuk tiap tiap link tetangga
			if !visited[neighbor] {
				newPath := append(path, neighbor)
				err := DLSAllPath(neighbor, target, depth-1, visited, newPath, visitedPath, allPaths, artikelDiperiksa)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func IDSAllPath(startURL string, targetURL string) ([][]string, int, int, error) {
	artikelDiperiksa := 0
	allPaths := [][]string{}
	startTime := time.Now()
	//melakukan pencarian untuk tiap kedalaman
	for depth := 0; time.Since(startTime).Seconds() <= 300; depth++ {
		//inisialisasi
		visited := make(map[string]bool)
		visitedPath := make(map[string]bool)
		path := []string{startURL}

		err := DLSAllPath(startURL, targetURL, depth, visited, path, visitedPath, &allPaths, &artikelDiperiksa)
		if err != nil {
			return nil, artikelDiperiksa, depth, err
		}

		if len(allPaths) > 0 {
			return allPaths, artikelDiperiksa, depth, nil
		}
	}

	return nil, artikelDiperiksa, 0, fmt.Errorf("tidak ditemukan path")
}
