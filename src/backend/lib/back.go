package backend

import (
	"fmt"
	"time"
)

type FormData struct {
	StartArticle string `json:"startArticle"`
	StartUrl     string `json:"startUrl"`
	EndArticle   string `json:"endArticle"`
	EndUrl       string `json:"endUrl"`
	Algoritma    string `json:"algoritma"`
}

type Node struct {
	Depth int
	Url   string
	Prev  *Node
}

func MainBackend(data FormData) ([][]string, int, int, time.Duration, error) {
	var paths [][]string
	var checkedArticle int
	var clickArticle int
	var excTime time.Duration
	var err error

	if data.Algoritma == "BFS" {
		fmt.Println("\nRute BFS")
		startTime := time.Now()
		var path []string
		path, checkedArticle, clickArticle, err = BFS(data.StartUrl, data.EndUrl)
		paths = append(paths, path)
		excTime = time.Since(startTime)
		if err != nil {
			return nil, 0, 0, 0, err
		}

	} else if data.Algoritma == "BFS All Path" {
		fmt.Println("\nAll Path BFS")
		startTime := time.Now()
		paths, checkedArticle, clickArticle, err = BFSAllPath(data.StartUrl, data.EndUrl)
		excTime = time.Since(startTime)
		if err != nil {
			return nil, 0, 0, 0, err
		}

	} else if data.Algoritma == "IDS" {
		fmt.Println("\nRute IDS")
		startTime := time.Now()
		var path []string
		path, checkedArticle, clickArticle, err = IDS(data.StartUrl, data.EndUrl)
		paths = append(paths, path)
		excTime = time.Since(startTime)
		if err != nil {
			return nil, 0, 0, 0, err
		}

	} else {
		fmt.Println("\nAll Path IDS")
		startTime := time.Now()
		paths, checkedArticle, clickArticle, err = IDSAllPath(data.StartUrl, data.EndUrl)
		excTime = time.Since(startTime)
		if err != nil {
			return nil, 0, 0, 0, err
		}
	}

	return paths, checkedArticle, clickArticle, excTime, nil
}
