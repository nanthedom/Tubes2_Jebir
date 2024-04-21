package backend

import(
	"fmt"
	"time"
)
func DLS(URL string, target string, depth int, visited map[string]bool, path []string) ([]string,int, error) {
	artikel_diperiksa := 1
	if depth == 0 && URL == target {
		return []string{URL}, artikel_diperiksa, nil
	}

	if depth > 0 {
		visited[URL] = true
		neighbors, err := scrapeLinks(URL)
		if err != nil {
			return nil, 0, err
		} 
			for _, neighbor := range neighbors {
				if visited[neighbor] == false {
					temp , count, err2:= DLS(neighbor, target, depth-1, visited, append(path, neighbor))
					if err2 != nil {
						return nil, 0, err2
					} 
					artikel_diperiksa += count
					if temp != nil {
						result := append([]string{URL}, temp...)
						return result, artikel_diperiksa, nil
					}
				}
			}
	}

	return nil, artikel_diperiksa, nil
}

func IDS(startURL string, targetURL string) ([]string, int,int, error){
	artikel_diperiksa := 0
	startTime := time.Now()
	for depth := 0; time.Since(startTime).Seconds() <= 300; depth++ {
		// fmt.Println("Kedalaman: ", depth)
		visited := make(map[string]bool)
		
		result, count, err := DLS(startURL, targetURL, depth, visited, []string{startURL})
		artikel_diperiksa += count
		if err != nil {
			return nil, artikel_diperiksa,0, err
		} 
		if result != nil {
			return result, artikel_diperiksa,len(result)-1, nil
		}
	}
	return nil, artikel_diperiksa,0, fmt.Errorf("not found")
}

