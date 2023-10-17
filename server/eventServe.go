package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"thamel/dist"
	"thamel/types"
)

func writeResponseToStreamEs(w http.ResponseWriter, events types.ResponseIE) {
	// Write the JSON response
	jsonData, err := json.Marshal(events)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Convert JSON to byte array
	byteArray := []byte(jsonData)
	w.Write(byteArray)
}

func EventServer(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Query().Get("slug")

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	if len(slug) == 0 {
		writeResponseToStreamEs(w, types.NullResponseIE())
		return
	}
	// Write the JSON response
	writeResponseToStreamEs(w, dist.IndivEvent(slug))
}
