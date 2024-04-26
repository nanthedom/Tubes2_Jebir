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
	// defer close(queue)

	var wg sync.WaitGroup
	var visitedMutex sync.Mutex

	found := false
	path := []string{}
	clickArticle := 0

	// isi queue dengan node start
	queue <- &Node{Url: start, Depth: 0}

	startNode := <-queue

	visited[startNode.Url] = true

	// scrape links dari start
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
			// queue = append(queue, &Node{Url: link, Prev: startNode, Depth: startNode.Depth + 1})
			queue <- &Node{Url: link, Prev: startNode, Depth: startNode.Depth + 1}
		}
	}

	// proses sisa queue
	for len(queue) > 0 && !found {
		currentNode := <-queue

		// print currentNode
		fmt.Println("Current Node URL: ", currentNode.Url)
		// print current checked article
		fmt.Println("Checked Article: ", checkedArticle)

		// sleep for 0.5 millisecond
		time.Sleep(10 * time.Millisecond)
		// time.Sleep(500 * time.Microsecond)

		wg.Add(1)

		go func(currentNode *Node) {
			defer wg.Done()

			// mutex cegah concurrent write dan read
			// defer visitedMutex.Unlock()
			visitedMutex.Lock()
			if visited[currentNode.Url] {
				visitedMutex.Unlock()
				return
			}
			visitedMutex.Unlock() // coba

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
				// mutex cegah concurrent write dan read
				visitedMutex.Lock()
				if !visited[link] {
					checkedArticle++
					if link == end {
						path = append(buildPathToTarget(currentNode), link)
						found = true
						// print
						fmt.Println("Found?", found)
						return
					}
					// queue = append(queue, &Node{Url: link, Prev: currentNode, Depth: currentNode.Depth + 1})
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
