package backend

import (
	"fmt"
	"log"
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

func MainBackend(data FormData) {
	if data.Algoritma == "BFS" {
		startTime := time.Now()
		paths, checkedArticle, clickArticle, err := BFS(data.StartUrl, data.EndUrl)
		excTime := time.Since(startTime)
		if err != nil {
			log.Fatal(err)
		}
		for _, path := range paths {
			fmt.Println("URL:", path)
		}
		fmt.Println("checked article: ", checkedArticle)
		fmt.Println("click article: ", clickArticle)
		fmt.Println("excecution time:", excTime)
	} else {
		startTime := time.Now()
		paths, checkedArticle, clickArticle := IDS(data.StartUrl, data.EndUrl)
		excTime := time.Since(startTime)
		for _, path := range paths {
			fmt.Println("URL:", path)
		}
		fmt.Println("checked article: ", checkedArticle)
		fmt.Println("click article: ", clickArticle)
		fmt.Println("excecution time:", excTime)
	}
}
