package main

import(
	// "fmt"
)
func dfs(node *Node, target string, depth int, visited map[string]bool, path []*Node) ([]*Node,int) {
	// fmt.Println("	url : ", node.URL)
	artikel_diperiksa := 1
	if depth == 0 && node.URL == target {
		return []*Node{node}, artikel_diperiksa
	}

	if depth > 0 {
		visited[node.URL] = true
		// neighbors, err := getNeighborsFromURL(node.URL)
		neighbors, err := getNeighborsFromURL(node.URL)
		if err != nil {
			return nil, 0
		} 
			for _, neighbor := range neighbors {
				if visited[neighbor.URL] == false {
					temp , count:= dfs(neighbor, target, depth-1, visited, append(path, neighbor))
					artikel_diperiksa += count
					if temp != nil {
						result := append([]*Node{node}, temp...)
						return result, artikel_diperiksa
					}
				}
			}
	}

	return nil, artikel_diperiksa
}

func iterativeDeepeningSearch(startURL string, targetURL string, depth_max int) ([]*Node, int,int){
	artikel_diperiksa := 0
	for depth := 0; depth <= depth_max; depth++ {
		// fmt.Println("Kedalaman: ", depth)
		startNode := &Node{Title: startURL, URL: startURL}
		visited := make(map[string]bool)
		
		result, count := dfs(startNode, targetURL, depth, visited, []*Node{startNode})
		artikel_diperiksa += count
		if result != nil {
			return result, artikel_diperiksa,count
		}
	}
	return nil, artikel_diperiksa,0
}

