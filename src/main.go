package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	backend "app/util/backend"
)

var formData backend.FormData

func handleInsert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(requestBody, &formData); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	fmt.Println("\nRequest body:", string(requestBody))
	fmt.Printf("Received data: %+v\n\n", formData)

	responseData := struct {
		Message      string `json:"message"`
		StartArticle string `json:"startArticle"`
		EndArticle   string `json:"endArticle"`
	}{
		Message:      "Received data successfully",
		StartArticle: formData.StartArticle,
		EndArticle:   formData.EndArticle,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseData)
}

func main() {
	http.HandleFunc("/", handleInsert)
	fmt.Println("Server listening on port 8000...")
	http.ListenAndServe(":8000", nil)
}
