// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"io"
// 	"log"
// 	"net/http"
// )

// const (
// 	openAIAPIURL = "https://api.openai.com/v1/chat/completions"
// 	chatgptToken = // Replace with your actual token
// )

// // // RequestPayload defines the structure of the request payload to OpenAI
// // type RequestPayload struct {
// // 	Model       string    `json:"model"`
// // 	Messages    []Message `json:"messages"`
// // 	Temperature float64   `json:"temperature"`
// // 	MaxTokens   int       `json:"max_tokens"`
// // }

// // // Message defines the structure of the message in the request payload
// // type Message struct {
// // 	Role    string `json:"role"`
// // 	Content string `json:"content"`
// // }

// // // ResponsePayload defines the structure of the response payload from OpenAI
// // type ResponsePayload struct {
// // 	Choices []Choice `json:"choices"`
// // }

// // // Choice defines the structure of a choice in the response payload
// // type Choice struct {
// // 	Message Message `json:"message"`
// // }

// //structure_of_payload(query) //json
// type RequestPayload struct {
// 	Model       string    `json:"model"`
// 	Messages    []Message `json:"messages"`
// 	Temperature float64   `json:"temperature"`
// 	MaxTokens   int       `json:"max_tokens"`
// }

//  //function_call_to_GPTAPI(payload){
//     //return response from chatgpt
//     //}

// func main() {

// 	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {

// 		var requestPayload RequestPayload
// 		// // Create the request to OpenAI
// 		jsonData, err := json.Marshal(requestPayload)
// 		if err != nil {
// 			http.Error(w, `{"error": "Failed to marshal request payload"}`, http.StatusInternalServerError)
// 			return
// 		}

// 		req, err := http.NewRequest(http.MethodPost, openAIAPIURL, bytes.NewBuffer(jsonData))
// 		if err != nil {
// 			http.Error(w, `{"error": "Failed to create request"}`, http.StatusInternalServerError)
// 			return
// 		}

// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("Authorization", "Bearer "+chatgptToken)

// 		client := &http.Client{}
// 		resp, err := client.Do(req)
// 		if err != nil {
// 			http.Error(w, `{"error": "Failed to make request to OpenAI"}`, http.StatusInternalServerError)
// 			return
// 		}
// 		defer resp.Body.Close()

// 		responseBody, err := io.ReadAll(resp.Body)

// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(responseBody)
// 	})

//		log.Println("Server starting on port 8080")
//		if err := http.ListenAndServe(":8080", nil); err != nil {
//			log.Fatalf("Server failed: %v", err)
//		}
//	}
//
// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const (
	openAIAPIURL = "https://api.openai.com/v1/chat/completions"
	chatgptToken =  // Replace with your actual token
)

// //Request
// RequestPayload defines the structure of the request payload to OpenAI
// type RequestPayload struct {
// 	Model       string    `json:"model"`
// 	Messages    []Message `json:"messages"`
// 	Temperature float64   `json:"temperature"`
// 	MaxTokens   int       `json:"max_tokens"`
// }

// Message defines the structure of the message in the request payload

type RequestPayload struct {
	Content string `json:"content"`
}

type OpenAIRequestPayload struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
	MaxTokens   int       `json:"max_tokens"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

/////Response

// ResponsePayload defines the structure of the response payload from OpenAI
type ResponsePayload struct {
	Choices []Choice `json:"choices"`
}

// Choice defines the structure of a choice in the response payload
type Choice struct {
	Message Message `json:"message"`
}

func main() {

	//get query from esp32
	//structure bnana h chatgpt format json me
	//chatgpt ko query bhejna h
	//fetch chatgpt response in json
	//print chatgpt query reply msg from json
	//send query reply to esp32

	//STEP1: get query from esp32
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request method: %s", r.Method) // Log the request method
		if r.Method != "POST" {
			http.Error(w, `{"error": "Invalid request method"}`, http.StatusMethodNotAllowed)
			return
		}

		// Read and log the request body
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, `{"error": "Failed to read query from esp32"}`, http.StatusInternalServerError)
			return
		}
		log.Printf("Received query from esp32: %s", string(bodyBytes)) // Log the raw request body

		//STEP 1: Complete

		//STEP 2: structure bnana h chatgpt format json me

		// Hardcoded values for OpenAI request
		openAIRequestPayload := OpenAIRequestPayload{
			Model:       "gpt-4",
			Messages:    []Message{{Role: "user", Content: string(bodyBytes)}},
			Temperature: 0.5,
			MaxTokens:   20,
		}

		// Create the request to OpenAI

		jsonData, err := json.Marshal(openAIRequestPayload)
		if err != nil {
			http.Error(w, `{"error": "Failed to marshal request payload"}`, http.StatusInternalServerError)
			return
		}

		req, err := http.NewRequest(http.MethodPost, openAIAPIURL, bytes.NewBuffer(jsonData))
		if err != nil {
			http.Error(w, `{"error": "Failed to create request"}`, http.StatusInternalServerError)
			return
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+chatgptToken)

		//STEP:3 chatgpt ko query bhejna h  //fetch chatgpt response in json
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, `{"error": "Failed to make request to OpenAI"}`, http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, `{"error": "Failed to read response body"}`, http.StatusInternalServerError)
			return
		}
		//STEP:4 print chatgpt query reply msg from json

		// Log the entire response from OpenAI for debugging
		log.Printf("Response from OpenAI: %s", responseBody)

		if resp.StatusCode != http.StatusOK {
			http.Error(w, `{"error": "OpenAI request failed"}`, resp.StatusCode)
			return
		}

		//STEP 5: send query reply to esp32

		// Forward the response back to the client
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
