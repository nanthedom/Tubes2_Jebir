package backend

import (
	"fmt"
	"sync"
	"time"
)

func BFSgorut(start, end string) ([]string, int, int, error) {
	if len(start) == 0 || len(end) == 0 {
		return nil, 0, 0, nil
	}

	checkedArticle := 1
	if start == end {
		return []string{start, end}, checkedArticle, 0, nil
	}
	visited := make(map[string]bool)
	queue := make(chan *Node, 40000000)

	var wg sync.WaitGroup
	var visitedMutex sync.Mutex

	found := false
	path := []string{}
	clickArticle := 0

	queue <- &Node{Url: start, Depth: 0}
	startNode := <-queue
	visited[startNode.Url] = true

	links, err := scrapeLinks(startNode.Url)
	if err != nil {
		return nil, checkedArticle, 0, err
	}

	clickArticle = startNode.Depth + 1
	for _, link := range links {
		if !visited[link] {
			checkedArticle++
			if link == end {
				path = append(buildPathToTarget(startNode), link)
				found = true
				break
			}
			queue <- &Node{Url: link, Prev: startNode, Depth: startNode.Depth + 1}
		}
	}

	for len(queue) > 0 && !found {
		currentNode := <-queue

		time.Sleep(5 * time.Millisecond)

		wg.Add(1)

		go func(currentNode *Node) {
			defer wg.Done()
			visitedMutex.Lock()
			if visited[currentNode.Url] {
				visitedMutex.Unlock()
				return
			}
			visitedMutex.Unlock()

			visitedMutex.Lock()
			visited[currentNode.Url] = true
			visitedMutex.Unlock()

			links, err := scrapeLinks(currentNode.Url)
			if err != nil {
				path = nil
				return
			}

			clickArticle = currentNode.Depth + 1
			for _, link := range links {
				visitedMutex.Lock()
				if !visited[link] {
					checkedArticle++
					if link == end {
						path = append(buildPathToTarget(currentNode), link)
						found = true
						return
					}
					queue <- &Node{Url: link, Prev: currentNode, Depth: currentNode.Depth + 1}
				}
				visitedMutex.Unlock()
			}
		}(currentNode)
	}

	if found {
		return path, checkedArticle, clickArticle, nil
	}

	wg.Wait()
	close(queue)

	return nil, checkedArticle, 0, fmt.Errorf("not found")
}
