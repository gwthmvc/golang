package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Slip represents the structure of the incoming JSON data from the API
type Slip struct {
	Advice string `json:"advice"`
}

type APIResponse struct {
	Slip Slip `json:"slip"`
}

func main() {
	url := "https://api.adviceslip.com/advice"

	client := http.Client{Timeout: time.Second * 5}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// DEBUG STEP: Read the entire response body as raw bytes
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading body: %v\n", err)
		return
	}

	// Print out what the server actually replied with
	fmt.Println("--- RAW SERVER RESPONSE ---")
	fmt.Println(string(bodyBytes))
	fmt.Println("---------------------------")

	// Now try to unmarshal the saved bytes into our struct
	var result APIResponse
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}

	fmt.Println(result.Slip.Advice)
}
