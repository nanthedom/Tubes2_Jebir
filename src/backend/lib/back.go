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
		fmt.Println("\nRute BFS")
		// coba bonus
		startTime := time.Now()
		pathsBonus, checked, click, err := BFSBonus(data.StartUrl, data.EndUrl)
		excTime = time.Since(startTime)
		if err != nil {
			return nil, 0, 0, 0, err
		}
		for _, path := range pathsBonus {
			for _, url := range path {
				fmt.Println(url)
			}
			fmt.Println()
		}
		fmt.Println(len(pathsBonus))
		fmt.Println(checked)
		fmt.Println(click)
		paths, checkedArticle, clickArticle, err = BFS(data.StartUrl, data.EndUrl)
		if err != nil {
			return nil, 0, 0, 0, err
		}
	} else {
		fmt.Println("\nRute IDS")
		startTime := time.Now()
		paths, checkedArticle, clickArticle, err = IDS(data.StartUrl, data.EndUrl)
		if err != nil {
			return nil, 0, 0, 0, err
		}
		excTime = time.Since(startTime)
	}

	return paths, checkedArticle, clickArticle, excTime, nil
}
