package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type APIResponse struct {
	Slip struct {
		Advice string `json:"advice"`
	} `json:"slip"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
		<h1>Welcome to your Go Web Server! 🚀</h1>
		<hr>
		<p>👉 Visit <a href="/advice">localhost:8080/advice</a> to trigger a concurrent background fetch!</p>
	`)
}

func adviceHandler(w http.ResponseWriter, r *http.Request) {
	ch := make(chan string)

	go func() {
		client := http.Client{Timeout: time.Second * 3}
		resp, err := client.Get("https://api.adviceslip.com/advice")
		if err != nil {
			ch <- fmt.Sprintf("Network connection failed: %v", err)
			return
		}
		defer resp.Body.Close()

		// 1. Read the raw text payload first to check what we got
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			ch <- "Failed to read the API response."
			return
		}

		// 2. Check if the server sent an HTTP error status (like 429 Too Many Requests or 502)
		if resp.StatusCode != http.StatusOK {
			ch <- fmt.Sprintf("API returned HTTP error status %d. Response snippet: %s", resp.StatusCode, string(bodyBytes[:50]))
			return
		}

		// 3. Try decoding the saved bytes
		var result APIResponse
		if err := json.Unmarshal(bodyBytes, &result); err != nil {
			// If it fails, print out the first 100 characters of the bad data to see what it is
			snippet := string(bodyBytes)
			if len(snippet) > 100 {
				snippet = snippet[:100]
			}
			ch <- fmt.Sprintf("JSON Decode Error. Server actually sent: <pre>%s...</pre>", snippet)
			return
		}

		ch <- result.Slip.Advice
	}()

	fetchedAdvice := <-ch

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h2>🔮 Response Status:</h2><blockquote>%s</blockquote>", fetchedAdvice)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/advice", adviceHandler)

	fmt.Println("Server is starting up on http://localhost:8080 ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
