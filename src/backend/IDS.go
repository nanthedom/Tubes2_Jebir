package backend

import(
	// "fmt"
	"time"
)
func DLS(URL string, target string, depth int, visited map[string]bool, path []string) ([]string,int) {
	// fmt.Println("	url : ", URL)
	artikel_diperiksa := 1
	if depth == 0 && URL == target {
		return []string{URL}, artikel_diperiksa
	}

	if depth > 0 {
		visited[URL] = true
		// neighbors, err := getNeighborsFromURL(node.URL)
		neighbors, err := scrapeLinks(URL)
		if err != nil {
			return nil, 0
		} 
			for _, neighbor := range neighbors {
				if visited[neighbor] == false {
					temp , count:= DLS(neighbor, target, depth-1, visited, append(path, neighbor))
					artikel_diperiksa += count
					if temp != nil {
						result := append([]string{URL}, temp...)
						return result, artikel_diperiksa
					}
				}
			}
	}

	return nil, artikel_diperiksa
}

func IDS(startURL string, targetURL string) ([]string, int,int){
	artikel_diperiksa := 0
	startTime := time.Now()
	for depth := 0; time.Since(startTime).Seconds() <= 300; depth++ {
		// fmt.Println("Kedalaman: ", depth)
		visited := make(map[string]bool)
		
		result, count := DLS(startURL, targetURL, depth, visited, []string{startURL})
		artikel_diperiksa += count
		if result != nil {
			return result, artikel_diperiksa,len(result)-1 
		}
	}
	return nil, artikel_diperiksa,0
}

