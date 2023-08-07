package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

type Temperature struct {
	City      string  `json:"city"`
	Timestamp string  `json:"timestamp"`
	Degree    float64 `json:"degree"`
}

func GetLocalTemperature(w http.ResponseWriter, r *http.Request) {
	// Setting the Content-Type to JSON since we'll return a JSON response
	w.Header().Set("Content-Type", "application/json")

    // Adding CORS headers
    w.Header().Set("Access-Control-Allow-Origin", "*")

	// Setting a custom header just for demonstration
	w.Header().Set("X-Custom-Header", "CustomValue")

	// Generate a mock temperature response
	temp := Temperature{
		City:      "San Francisco",
		Timestamp: time.Now().Format(time.RFC3339),
		Degree:    22.5, // This can be fetched from a real service in a real-world scenario
	}

	// Convert the temperature struct to JSON
	data, err := json.Marshal(temp)
	if err != nil {
		// If there's an error, set the status to 500 Internal Server Error and return
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to generate temperature data."))
		return
	}

	// For demonstration, writing a snippet of HTML before the actual JSON data.
	// htmlData := []byte("<h1>Temperature Data:</h1>\n")
	// w.Write(htmlData)

	// Write the JSON data as the response body
	w.Write(data)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/temperature", GetLocalTemperature).Methods("GET")

	http.Handle("/", r)
	log.Println("Server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}