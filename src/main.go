package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	backend "app/util/backend"

	"github.com/rs/cors"
)

func handleInsert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var formData backend.FormData
	if err := json.NewDecoder(r.Body).Decode(&formData); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	fmt.Printf("\nReceived data:\n")
	fmt.Printf("Start Article: %s\n", formData.StartArticle)
	fmt.Printf("Start Url: %s\n", formData.StartUrl)
	fmt.Printf("End Article: %s\n", formData.EndArticle)
	fmt.Printf("End Url: %s\n", formData.EndUrl)
	fmt.Printf("Algoritma: %s\n", formData.Algoritma)

	paths, checkedArticle, clickArticle, excTime, err := backend.MainBackend(formData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseData := struct {
		Message        string   `json:"message"`
		StartArticle   string   `json:"startArticle"`
		EndArticle     string   `json:"endArticle"`
		StartUrl       string   `json:"startUrl"`
		EndUrl         string   `json:"endUrl"`
		Algoritma      string   `json:"algoritma"`
		Paths          []string `json:"paths"`
		CheckedArticle int      `json:"checkedArticle"`
		ClickArticle   int      `json:"clickArticle"`
		ExcTime        string   `json:"excTime"`
	}{
		Message:        "Received data successfully",
		StartArticle:   formData.StartArticle,
		EndArticle:     formData.EndArticle,
		StartUrl:       formData.StartUrl,
		EndUrl:         formData.EndUrl,
		Algoritma:      formData.Algoritma,
		Paths:          paths,
		CheckedArticle: checkedArticle,
		ClickArticle:   clickArticle,
		ExcTime:        excTime.String(),
	}

	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		http.Error(w, "Failed to encode response data", http.StatusInternalServerError)
		return
	}
}

func main() {
	// Create a new CORS handler allowing requests from localhost:3000
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	// Wrap the default ServeMux with the CORS handler
	handler := corsHandler.Handler(http.DefaultServeMux)
	http.HandleFunc("/", handleInsert)

	fmt.Println("Server listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", handler))

}
