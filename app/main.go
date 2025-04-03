package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Timestamp string `json:"timestamp"`
	IP        string `json:"ip"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Response{
		Timestamp: time.Now().Format(time.RFC3339),
		IP:        r.RemoteAddr,
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
