package backend

import (
	"fmt"
	"time"
)

func DLS(URL string, target string, depth int, visited map[string]bool) ([]string, int, error) {
	artikel_diperiksa := 0
	if depth == 0 && URL == target { //mengembalikan url jika merupakan link target
		artikel_diperiksa++
		return []string{URL}, artikel_diperiksa, nil
	}

	if depth == 0 { //menambahkan artikel diperiksa saat kedalaman == 0 (artikel yang baru)
		artikel_diperiksa++
	}

	if depth > 0 {
		visited[URL] = true
		neighbors, err := scrapeLinks(URL) //mencari link tetangga
		if err != nil {
			return nil, 0, err
		}
		for _, neighbor := range neighbors { //menerapkan DLS untuk tiap tiap link tetangga
			if !visited[neighbor] {
				temp, count, err2 := DLS(neighbor, target, depth-1, visited)
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

func IDS(startURL string, targetURL string) ([]string, int, int, error) {
	//menangani kasus nil
	if len(startURL) == 0 || len(targetURL) == 0 {
		return nil, 0, 0, nil
	}

	artikel_diperiksa := 0
	startTime := time.Now()
	//melakukan pencarian untuk tiap kedalaman
	for depth := 0; time.Since(startTime).Seconds() <= 300; depth++ {
		visited := make(map[string]bool)

		result, count, err := DLS(startURL, targetURL, depth, visited)
		artikel_diperiksa += count //menambahkan jumlah artikel yang diperiksa
		if err != nil {
			return nil, artikel_diperiksa, 0, err
		}
		if result != nil {
			return result, artikel_diperiksa, len(result) - 1, nil
		}
	}
	return nil, artikel_diperiksa, 0, fmt.Errorf("tidak ditemukan path")
}
