package backend

import (
	"fmt"
	"strings"
)

func BFSAllPath(start, end string) ([][]string, int, int, error) {
	//menangani kasus nil
	if len(start) == 0 || len(end) == 0 {
		return nil, 0, 0, nil
	}

	checkedArticle := 1
	clickedArticle := 1
	var paths [][]string
	if start == end { //jika start == end langsung return
		path := []string{start}
		return append(paths, path), checkedArticle, clickedArticle, nil
	}

	visited := make(map[string]bool)
	visitedPaths := make(map[string]bool)
	queue := []*Node{{Url: start, Depth: 0}}
	depth := 9999
	for len(queue) > 0 { //jika queue belum 0 maka loop terus
		currentNode := queue[0]
		queue = queue[1:]

		if !visited[currentNode.Url] {
			visited[currentNode.Url] = true
			if currentNode.Depth < depth {
				links, err := scrapeLinks(currentNode.Url) //mencari link tetangga
				if err != nil {
					return nil, checkedArticle, 0, err
				}

				for _, link := range links {
					if !visited[link] {
						checkedArticle++ // tambahkan link yang di periksa bila belum ada di visited
						if link == end {
							newPath := append(buildPathToTarget(currentNode), link)
							pathString := strings.Join(newPath, "->")
							if !visitedPaths[pathString] {
								visitedPaths[pathString] = true
								depth = currentNode.Depth + 1
								paths = append(paths, newPath) //mengumpulkan semua path
							}
						}
						if currentNode.Depth+1 < depth {
							//menambahkan link tetangga di queue
							queue = append(queue, &Node{Url: link, Prev: currentNode, Depth: currentNode.Depth + 1})
						}
					}
				}
			}
		}
	}
	if len(paths) == 0 {
		return nil, checkedArticle, 0, fmt.Errorf("not found")
	}
	return paths, checkedArticle, depth, nil
}
