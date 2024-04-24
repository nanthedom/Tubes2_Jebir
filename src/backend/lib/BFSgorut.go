package backend

// import (
// 	"fmt"
// 	"sync"
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
// 	// wg.Add(1)

// 	// Create a mutex to prevent race conditions
// 	var mutex sync.Mutex

// 	var mutex2 sync.Mutex

// 	// go func() {
// 	// defer wg.Done()
// 	queue <- &Node{Url: start, Depth: 0}
// 	// }()

// 	for len(queue) > 0 {
// 		// select {
// 		currentNode := <-queue

// 		// print currentNode
// 		fmt.Println("Current Node: ", currentNode.Url)
// 		// print current checked article
// 		fmt.Println("Checked Article: ", checkedArticle)

// 		// Process current node concurrently

// 		wg.Add(1)

// 		go func(node *Node) {
// 			// Defer the completion of the current goroutine
// 			defer wg.Done()

// 			// Lock the mutex to prevent race conditions
// 			mutex.Lock()
// 			if visited[node.Url] {
// 				mutex.Unlock()
// 				return
// 			}
// 			visited[node.Url] = true
// 			mutex.Unlock()

// 			if node.Url == end {
// 				// Build the 	path to the target
// 				path := buildPathToTarget(node)
// 				// Signal to the main goroutine that we found the target
// 				close(queue)
// 				// Return the path and other information
// 				returnResult(path, checkedArticle, node.Depth)
// 				return
// 			}

// 			// Scrape links for the current node
// 			links, err := scrapeLinks(node.Url)
// 			if err != nil {
// 				// Handle error
// 				close(queue)
// 				returnResult(nil, checkedArticle, 0)
// 				return
// 			}

// 			// Increment checked article count
// 			mutex2.Lock()
// 			checkedArticle++
// 			mutex2.Unlock()

// 			// Add nodes to the queue
// 			for _, link := range links {
// 				// if link == end, we found the target, return with the path to currentNode and append this link node to the path
// 				if link == end {
// 					close(queue)
// 					returnResult(append(buildPathToTarget(node), link), checkedArticle, node.Depth+1)
// 					return
// 				}

// 				queue <- &Node{Url: link, Prev: node, Depth: node.Depth + 1}
// 			}
// 		}(currentNode)

// 		// // Timeout handling
// 		// case <-time.After(5 * time.Second):
// 		// 	close(queue)
// 		// 	return nil, checkedArticle, 0, fmt.Errorf("timeout")
// 		// }
// 	}

// 	// Wait for all goroutines to finish
// 	wg.Wait()

// 	// Return an error if the target is not found
// 	return nil, checkedArticle, 0, fmt.Errorf("not found")
// }

// func returnResult(path []string, checkedArticle, clickArticle int) {
// 	// // Handle the result here, e.g., print or return
// 	// fmt.Println("Path:", path)
// 	// fmt.Println("Checked Articles:", checkedArticle)
// 	// fmt.Println("Click Articles:", clickArticle)
// }






// ATAU -----------------
// package backend

// import (
// 	"fmt"
// 	"sync"
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
// 	// wg.Add(1)

// 	// Create a mutex to prevent race conditions
// 	var mutex sync.Mutex

// 	var mutex2 sync.Mutex

// 	// go func() {
// 	// defer wg.Done()
// 	queue <- &Node{Url: start, Depth: 0}
// 	// }()

// 	// boolean found to check if the target is found
// 	found := false
// 	path := []string{}
// 	nodeDepth := 0

// 	for len(queue) > 0 && !found {
// 		// select {
// 		currentNode := <-queue

// 		// print currentNode
// 		fmt.Println("Current Node: ", currentNode.Url)
// 		// print current checked article
// 		fmt.Println("Checked Article: ", checkedArticle)

// 		// Process current node concurrently

// 		wg.Add(1)

// 		go func(node *Node) {
// 			// Defer the completion of the current goroutine
// 			defer wg.Done()

// 			// Lock the mutex to prevent race conditions
// 			mutex.Lock()
// 			if visited[node.Url] {
// 				mutex.Unlock()
// 				return
// 			}
// 			visited[node.Url] = true
// 			mutex.Unlock()

// 			if node.Url == end {
// 				// append path with buildpathtotarget
// 				path = append(path, buildPathToTarget(node)...)
// 				nodeDepth = node.Depth
// 				// Signal to the main goroutine that we found the target
// 				close(queue)
// 				// Return the path and other information
// 				// returnResult(path, checkedArticle, node.Depth)
// 				found = true
// 				return
// 			}

// 			// Scrape links for the current node
// 			links, err := scrapeLinks(node.Url)
// 			if err != nil {
// 				// Handle error
// 				close(queue)
// 				// returnResult(nil, checkedArticle, 0)
// 				return
// 			}

// 			// Increment checked article count
// 			mutex2.Lock()
// 			checkedArticle++
// 			mutex2.Unlock()

// 			// Add nodes to the queue
// 			for _, link := range links {
// 				// if link == end, we found the target, return with the path to currentNode and append this link node to the path
// 				if link == end {
// 					close(queue)
// 					path = append(buildPathToTarget(node), link)
// 					nodeDepth = node.Depth + 1
// 					// returnResult(append(buildPathToTarget(node), link), checkedArticle, node.Depth+1)
// 					found = true
// 					return
// 				}

// 				queue <- &Node{Url: link, Prev: node, Depth: node.Depth + 1}
// 			}
// 		}(currentNode)

// 		// // Timeout handling
// 		// case <-time.After(5 * time.Second):
// 		// 	close(queue)
// 		// 	return nil, checkedArticle, 0, fmt.Errorf("timeout")
// 		// }
// 	}

// 	// if found return path, checkedArticle, node.depth, nil
// 	if found {
// 		return path, checkedArticle, nodeDepth, nil
// 	}

// 	// Wait for all goroutines to finish
// 	wg.Wait()

// 	// Return an error if the target is not found
// 	return nil, checkedArticle, 0, fmt.Errorf("not found")
// }
