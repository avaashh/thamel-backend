package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"thamel/server"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]int{"status": 200}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	http.HandleFunc("/list", server.ListServer)
	// http.HandleFunc("/tags", server.TagsServer)
	http.HandleFunc("/event", server.EventServer)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
