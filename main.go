package main

import (
	"fmt"
	"net/http"
	"thamel/server"
)

func main() {
	http.HandleFunc("/list", server.ListServer)
	// http.HandleFunc("/tags", server.TagsServer)
	http.HandleFunc("/event", server.EventServer)

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
