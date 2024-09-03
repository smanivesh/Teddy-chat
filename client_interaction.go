package client_interaction
import main.go

import (
	"main"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

//get query from esp32
//OpenAI function call via main function
//send response to esp32
func client_interaction() {
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request method: %s", r.Method) // Log the request method
		if r.Method != "POST" {
			http.Error(w, `{"error": "Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		// Read query from user in json
		query, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, `{"error": "Failed to read request body"}`, http.StatusInternalServerError)
			return
		}
		log.Printf("Received request body from ESP32: %s", query.messages.content) // Log the raw request body
		

		// Call main function




		//send response back to esp32

		w.Header().Set("Content-Type", "application/json")
		w.Write(responseBody)
		// w.WriteHeader(http.StatusCreated)
		// w.Write([]byte("blabla\n"))
	})

	log.Println("Server starting on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
