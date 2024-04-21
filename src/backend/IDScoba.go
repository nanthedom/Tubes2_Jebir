package backend

import "time"

func IDS(start, end string) ([]string, int, int) {
	var path []string
	var checkedArticle, clickArticle int
	startTime := time.Now()

	for depth := 0; time.Since(startTime).Seconds() <= 300; depth++ {
		path, checkedArticle, clickArticle = DLS(start, end, depth)
		if path != nil {
			return path, checkedArticle, clickArticle
		}
	}

	return nil, checkedArticle, clickArticle
}

func DLS(start, end string, depthLimit int) ([]string, int, int) {
	checkedArticle := 0
	visited := make(map[string]bool)
	stack := []*Node{{Url: start, Depth: 0}}

	for len(stack) > 0 {
		currentNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		checkedArticle++
		clickArticle := currentNode.Depth
		if currentNode.Url == end {
			return buildPathToTarget(currentNode), checkedArticle, clickArticle
		}

		if currentNode.Depth < depthLimit {
			visited[currentNode.Url] = true
			links, err := scrapeLinks(currentNode.Url)
			if err != nil {
				return nil, checkedArticle, 0
			}
			for _, link := range links {
				if !visited[link] {
					stack = append(stack, &Node{Url: link, Prev: currentNode, Depth: currentNode.Depth + 1})
				}
			}
		}
	}

	return nil, checkedArticle, 0
}
