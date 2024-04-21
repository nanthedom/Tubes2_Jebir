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

func MainBackend(data FormData) ([]string, int, int, time.Duration, error) {
	var paths []string
	var checkedArticle int
	var clickArticle int
	var excTime time.Duration
	var err error

	if data.Algoritma == "BFS" {
		startTime := time.Now()
		paths, checkedArticle, clickArticle, err = BFS(data.StartUrl, data.EndUrl)
		excTime = time.Since(startTime)
		if err != nil {
			return nil, 0, 0, 0, err
		}
	} else {
		startTime := time.Now()
		paths, checkedArticle, clickArticle, err = IDS(data.StartUrl, data.EndUrl)
		if err != nil {
			return nil, 0, 0, 0, err
		}
		excTime = time.Since(startTime)
	}

	for _, path := range paths {
		fmt.Println("URL:", path)
	}

	return paths, checkedArticle, clickArticle, excTime, nil
}
