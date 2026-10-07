package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Slip struct {
	Advice string `json:"advice"`
}

type APIResponse struct {
	Slip Slip `json:"slip"`
}

// fetchAdvice is our worker function. It runs in the background.
// It accepts an ID number and a channel to send the finished text back to.
func fetchAdvice(workerID int, ch chan string) {
	url := "https://api.adviceslip.com/advice"
	client := http.Client{Timeout: time.Second * 5}

	fmt.Printf("[Worker %d] Starting API request...\n", workerID)

	resp, err := client.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("[Worker %d] Failed: %v", workerID, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		ch <- fmt.Sprintf("[Worker %d] API request failed: %s", workerID, resp.Status)
		return
	}

	var result APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		ch <- fmt.Sprintf("[Worker %d] JSON Error: %v", workerID, err)
		return
	}

	// Send the data back through the channel pipe
	ch <- fmt.Sprintf("[Worker %d] Success! Advice: %s", workerID, result.Slip.Advice)
}

func main() {
	startTime := time.Now()

	// 1. Create a channel that carries 'string' data
	adviceChannel := make(chan string)

	// 2. Launch 3 background tasks simultaneously using the 'go' keyword
	fmt.Println("Launching 3 workers at the exact same time...")
	for i := 1; i <= 3; i++ {
		go fetchAdvice(i, adviceChannel)
	}

	// 3. Wait for and collect the results from the channel
	// The main function will pause here until data arrives in the pipe
	fmt.Println("Waiting for workers to report back...")
	for i := 1; i <= 3; i++ {
		result := <-adviceChannel
		fmt.Println(result)
	}

	// 4. Calculate how long it took
	fmt.Printf("\nAll workers finished! Total time elapsed: %v\n", time.Since(startTime))
}
