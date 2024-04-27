package backend

import (
	"fmt"
	"sync"
	"time"
)

func DLSgorut(URL string, target string, depth int, visited map[string]bool, ch chan []string) (int, error) {

	// print URL
	fmt.Println("URL: ", URL)

	artikel_diperiksa := 0
	if depth == 0 && URL == target {
		artikel_diperiksa++
		ch <- []string{URL}
		// print
		fmt.Println("MASUK AAAAAAAAAAAAAAAAAAAAAAAAA")
		return artikel_diperiksa, nil
	}

	if depth == 0 {
		artikel_diperiksa++
	}

	if depth > 0 {
		visited[URL] = true
		neighbors, err := scrapeLinks(URL)
		if err != nil {
			// print error
			fmt.Println("ERROR 3")
			return 0, err
		}

		wg := sync.WaitGroup{}
		// wg.Add(len(neighbors))
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				// print
				fmt.Println("TIME 2")
				// sleep for 500 millisecond
				time.Sleep(10 * time.Microsecond)
				wg.Add(1)
				go func(neighbor string) {
					defer wg.Done()
					count, err2 := DLSgorut(neighbor, target, depth-1, make(map[string]bool), ch)
					artikel_diperiksa += count

					// print artikel diperiksa
					fmt.Println("Artikel diperiksa: ", artikel_diperiksa)

					if err2 != nil {
						// handle error (optional)
						// print error
						fmt.Println("ERROR 2")
						return
					}
				}(neighbor)
			}
		}
		wg.Wait()
	}

	return artikel_diperiksa, nil
}

func IDSgorut(startURL string, targetURL string) ([]string, int, int, error) {
	if len(startURL) == 0 || len(targetURL) == 0 {
		return nil, 0, 0, nil
	}

	artikel_diperiksa := 0
	startTime := time.Now()
	ch := make(chan []string, 16000000)

	for depth := 0; time.Since(startTime).Seconds() <= 300; depth++ {

		// print depth
		fmt.Println("Depth: ", depth)

		visited := make(map[string]bool)
		wg := sync.WaitGroup{}

		// print
		fmt.Println("TIME 1")
		// sleep for 500 millisecond
		time.Sleep(1 * time.Millisecond)

		wg.Add(1)
		go func() {
			defer wg.Done()
			// print
			fmt.Println("MASIH JALAN")
			count, err := DLSgorut(startURL, targetURL, depth, visited, ch)
			artikel_diperiksa += count
			if err != nil {
				// Handle the error (e.g., log it or return an error)
				// print error
				fmt.Println("ERROR 1")
				return
			}
		}()
		wg.Wait()

		select {
		case path := <-ch:
			return path, artikel_diperiksa, len(path) - 1, nil
		default:
			// continue loop if no path received yet
		}
	}
	return nil, artikel_diperiksa, 0, fmt.Errorf("not found")
}
