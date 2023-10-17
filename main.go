package handler

import (
	"encoding/json"
	"net/http"
	"thamel/server"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		status(w, r)
	case "/list":
		server.ListServer(w, r)
	case "/event":
		server.EventServer(w, r)
	default:
		noImplementation(w, r)
	}
}

func status(w http.ResponseWriter, r *http.Request) {
	response := map[string]int{"status": 200}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func noImplementation(w http.ResponseWriter, r *http.Request) {
	response := map[string]int{"status": 501}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(response)
}
