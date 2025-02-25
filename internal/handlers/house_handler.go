package handlers

import (
	"net/http"
)

func HouseHandler(w http.ResponseWriter, r *http.Request) {
	// Implement house-related HTTP handling logic
	w.Write([]byte("House handler"))
} 