package backend

import "fmt"

func BFS(start, end string) ([]string, int, int, error) {
	if len(start) == 0 || len(end) == 0 {
		return nil, 0, 0, nil
	}

	checkedArticle := 1
	if start == end {
		return []string{start, end}, checkedArticle, 0, nil
	}
	visited := make(map[string]bool)
	queue := []*Node{{Url: start, Depth: 0}}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		if visited[currentNode.Url] {
			continue
		}

		visited[currentNode.Url] = true

		links, err := scrapeLinks(currentNode.Url)
		if err != nil {
			return nil, checkedArticle, 0, err
		}

		clickArticle := currentNode.Depth + 1
		for _, link := range links {
			if !visited[link] {
				checkedArticle++
				if link == end {
					return append(buildPathToTarget(currentNode), link), checkedArticle, clickArticle, nil
				}
				queue = append(queue, &Node{Url: link, Prev: currentNode, Depth: currentNode.Depth + 1})
			}
		}
	}
	return nil, checkedArticle, 0, fmt.Errorf("not found")
}
