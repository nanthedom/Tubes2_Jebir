package backend

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func BFSAllPathGorut(start, end string) ([][]string, int, int, error) {
	if len(start) == 0 || len(end) == 0 {
		return nil, 0, 0, nil
	}

	// found := false
	// path := []string{}
	depth := 9999
	checkedArticle := 1
	clickedArticle := 1
	var paths [][]string
	if start == end {
		path := []string{start, end}
		return append(paths, path), checkedArticle, clickedArticle, nil
	}

	visited := make(map[string]bool)
	visitedPaths := make(map[string]bool)
	// queue := []*Node{{Url: start, Depth: 0}}
	queue := make(chan *Node, 40000000)

	var wg sync.WaitGroup
	var visitedMutex sync.Mutex
	var visitedPathsMutex sync.Mutex

	// isi queue dengan node start
	queue <- &Node{Url: start, Depth: 0}

	startNode := <-queue

	visited[startNode.Url] = true

	// scrape links dari start
	if startNode.Depth < depth {
		links, err := scrapeLinks(startNode.Url)
		if err != nil {
			return nil, checkedArticle, 0, err
		}

		for _, link := range links {
			if !visited[link] {
				checkedArticle++
				if link == end {
					newPath := append(buildPathToTarget(startNode), link)
					pathString := strings.Join(newPath, "->")
					if !visitedPaths[pathString] {
						visitedPaths[pathString] = true
						depth = startNode.Depth + 1
						paths = append(paths, newPath)
					}
				}
				if startNode.Depth+1 < depth {
					queue <- &Node{Url: link, Prev: startNode, Depth: startNode.Depth + 1}
				}
			}
		}
	}

	for len(queue) > 0 {
		// currentNode := queue[0]
		// queue = queue[1:]
		currentNode := <-queue

		// print currentNode
		fmt.Println("Current Node URL: ", currentNode.Url)
		// print current checked article
		fmt.Println("Checked Article: ", checkedArticle)

		time.Sleep(25 * time.Millisecond)

		wg.Add(1)

		go func(currentNode *Node) {
			defer wg.Done()

			if !visited[currentNode.Url] {
				visitedMutex.Lock()
				visited[currentNode.Url] = true
				visitedMutex.Unlock()
				if currentNode.Depth < depth {
					links, err := scrapeLinks(currentNode.Url)
					if err != nil {
						paths = nil
						clickedArticle = 0
						return
					}

					for _, link := range links {
						// visitedMutex.Lock()
						if !visited[link] {
							checkedArticle++
							if link == end {
								newPath := append(buildPathToTarget(currentNode), link)
								pathString := strings.Join(newPath, "->")
								if !visitedPaths[pathString] {
									visitedPathsMutex.Lock()
									visitedPaths[pathString] = true
									visitedPathsMutex.Unlock()
									depth = currentNode.Depth + 1
									paths = append(paths, newPath)
								}
								return
							}
							if currentNode.Depth+1 < depth {
								// queue = append(queue, &Node{Url: link, Prev: currentNode, Depth: currentNode.Depth + 1})
								queue <- &Node{Url: link, Prev: currentNode, Depth: currentNode.Depth + 1}
							}
						}
						// visitedMutex.Unlock()
					}
				}
			}
		}(currentNode)

	}
	if len(paths) == 0 {
		return nil, checkedArticle, 0, fmt.Errorf("not found")
	}

	wg.Wait()
	close(queue)

	return paths, checkedArticle, depth, nil
}
