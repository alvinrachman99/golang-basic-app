package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struktur untuk response JSON
type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// Handler untuk endpoint /
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "Welcome to the Go HTTP Server!",
		Status:  200,
	}
	json.NewEncoder(w).Encode(response)
}

// Handler untuk endpoint /hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // Ambil query parameter "name"
	if name == "" {
		name = "World"
	}
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: fmt.Sprintf("Hello, %s!", name),
		Status:  200,
	}
	json.NewEncoder(w).Encode(response)
}

// Main function
func main() {
	// Daftarkan handler
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)

	// Jalankan server di port 8080
	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
