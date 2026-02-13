package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	
	"github.com/kodanda2301/orchestrix/pkg/api" // Import the shared types
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	// In the future, we will read real CPU stats here.
	// For now, we return a mock status.
	status := api.HealthStatus{
		Status:    "ready",
		CPUUsage:  12.5,
		RAMFree:   8192,
		ActiveJob: false,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func main() {
	fmt.Println("🔧 Worker Node starting on port 8081...")

	http.HandleFunc("/health", healthHandler) // <--- New endpoint

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}