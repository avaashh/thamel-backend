package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"thamel/dist"
	"thamel/types"
)

func writeResponseToStreamLs(w http.ResponseWriter, events types.Response25) {
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

func ListServer(w http.ResponseWriter, r *http.Request) {

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	sdate := r.URL.Query().Get("sdate")

	if len(sdate) == 0 {
		writeResponseToStreamLs(w, dist.TodayDailyEvents())
		return
	}
	// Write the JSON response
	writeResponseToStreamLs(w, dist.LiveEvents(sdate))
}
