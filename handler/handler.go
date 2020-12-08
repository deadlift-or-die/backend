package handler

import "net/http"

// RegisterHandlers registers the http handlers
func RegisterHandlers() {
	http.HandleFunc("/ping", pingHandler)
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
