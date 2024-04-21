package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	backend "app/util/backend"
)

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

	// Membaca body request
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var formData backend.FormData
	if err := json.Unmarshal(requestBody, &formData); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Menampilkan data yang diterima dari request
	fmt.Printf("\nReceived data:\n")
	fmt.Printf("Start Article: %s\n", formData.StartArticle)
	fmt.Printf("Start Url: %s\n", formData.StartUrl)
	fmt.Printf("End Article: %s\n", formData.EndArticle)
	fmt.Printf("End Url: %s\n", formData.EndUrl)
	fmt.Printf("Algoritma: %s\n", formData.Algoritma)

	// Memanggil fungsi MainBackend untuk memproses data
	paths, checkedArticle, clickArticle, excTime, err := backend.MainBackend(formData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mengirimkan hasil kembali sebagai JSON response
	responseData := map[string]interface{}{
		"message":        "Received data successfully",
		"startArticle":   formData.StartArticle,
		"endArticle":     formData.EndArticle,
		"startUrl":       formData.StartUrl,
		"endUrl":         formData.EndUrl,
		"algoritma":      formData.Algoritma,
		"paths":          paths,
		"checkedArticle": checkedArticle,
		"clickArticle":   clickArticle,
		"excTime":        excTime.String(),
	}
	json.NewEncoder(w).Encode(responseData)
}

func main() {
	http.HandleFunc("/", handleInsert)
	fmt.Println("Server listening on port 8000...")
	http.ListenAndServe(":8000", nil)
}
