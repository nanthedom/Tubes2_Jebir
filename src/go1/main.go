package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	startURL := "https://en.wikipedia.org/wiki/Joko_Widodo"
	targetURL := "https://en.wikipedia.org/wiki/Erina_Gudono"
	// startURL := "a"
	// targetURL := "g"
	
	var pilihan int;
	fmt.Println("1. BFS")
	fmt.Println("2. IDS")
	fmt.Print("Masukkan pilihan Anda: ")
	fmt.Scanf("%d", &pilihan)

	if pilihan == 1{

		startTime := time.Now()
	
		fmt.Println("Mencari path . . .")
		path, articlesChecked, articlesTraversed, err := BFS(startURL, targetURL)
	
		elapsed := time.Since(startTime)
	
		if err != nil {
			log.Fatal(err)
		}
	
		fmt.Println("\nPath dari", startURL, "ke", targetURL, "adalah:")
		for i, node := range path {
			fmt.Println(i+1, ".", node.Title, ":", node.URL)
		}
		fmt.Println("\nArtikel diperiksa:", articlesChecked)
		fmt.Println("Artikel dikunjungi:", articlesTraversed)
		fmt.Println("Waktu eksekusi:", elapsed)
	} else if pilihan ==2 {

		maxDepth := 5              // Maximum depth for iterative deepening search
		startTime := time.Now()
		fmt.Println("Mencari path . . .")
		result , diperiksa, dilalui := iterativeDeepeningSearch(startURL, targetURL, maxDepth)
		endTime := time.Since(startTime)
		
		if result == nil {
			fmt.Println("No path found within the depth limit")
		} else {
			fmt.Println("\nArtikel diperiksa sebanyak: ", diperiksa)
			fmt.Println("\nArtikel dilalui sebanyak: ", dilalui)
			fmt.Println("\nPath found:")
			for _, node := range result {
				fmt.Println(node.URL)
			}

			fmt.Println("\nWaktu eksekusi: ", endTime)
		}
	} else {
		fmt.Println("input salah")
	}
}