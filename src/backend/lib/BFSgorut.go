package backend

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func BFS(start, end string) ([]string, int, int, error) {
// 	if len(start) == 0 || len(end) == 0 {
// 		return nil, 0, 0, nil
// 	}

// 	checkedArticle := 1
// 	if start == end {
// 		return []string{start, end}, checkedArticle, 0, nil
// 	}

// 	visited := make(map[string]bool)
// 	queue := make(chan *Node)
// 	defer close(queue)

// 	// Create a wait group to wait for all goroutines to finish
// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	// Create a mutex to prevent race conditions
// 	var mutex sync.Mutex

// 	go func() {
// 		defer wg.Done()
// 		queue <- &Node{Url: start, Depth: 0}
// 	}()

// 	for {
// 		select {
// 		case currentNode := <-queue:
// 			// Process current node concurrently
// 			go func(node *Node) {
// 				// Defer the completion of the current goroutine
// 				defer wg.Done()

// 				// Lock the mutex to prevent race conditions
// 				mutex.Lock()
// 				if visited[node.Url] {
// 					mutex.Unlock()
// 					return
// 				}
// 				visited[node.Url] = true
// 				mutex.Unlock()

// 				if node.Url == end {
// 					// Build the 	path to the target
// 					path := buildPathToTarget(node)
// 					// Signal to the main goroutine that we found the target
// 					close(queue)
// 					// Return the path and other information
// 					returnResult(path, checkedArticle, node.Depth)
// 					return
// 				}

// 				// Scrape links for the current node
// 				links, err := scrapeLinks(node.Url)
// 				if err != nil {
// 					// Handle error
// 					close(queue)
// 					returnResult(nil, checkedArticle, 0)
// 					return
// 				}

// 				// Increment checked article count
// 				mutex.Lock()
// 				checkedArticle++
// 				mutex.Unlock()

// 				// Add nodes to the queue
// 				for _, link := range links {
// 					// if link == end, we found the target, return with the path to currentNode and append this link node to the path
// 					if link == end {
// 						close(queue)
// 						returnResult(append(buildPathToTarget(node), link), checkedArticle, node.Depth+1)
// 						return
// 					}

// 					queue <- &Node{Url: link, Prev: node, Depth: node.Depth + 1}
// 				}
// 			}(currentNode)

// 		// Timeout handling
// 		case <-time.After(5 * time.Second):
// 			close(queue)
// 			return nil, checkedArticle, 0, fmt.Errorf("timeout")
// 		}
// 	}

// 	// Wait for all goroutines to finish
// 	wg.Wait()

// 	// Return an error if the target is not found
// 	return nil, checkedArticle, 0, fmt.Errorf("not found")
// }

// func returnResult(path []string, checkedArticle, clickArticle int) {
// 	// Handle the result here, e.g., print or return
// 	fmt.Println("Path:", path)
// 	fmt.Println("Checked Articles:", checkedArticle)
// 	fmt.Println("Click Articles:", clickArticle)
// }
