# Learning Go: From Basics to Concurrent Web Servers

A hands-on repository tracking my journey learning **Go (Golang)**. This project progresses from basic console printing to fetching data concurrently from an external API, and finally exposing that data through a local web server.

## 🚀 Project Evolution *so far*

This repository is structured around three key milestones in my Go learning path:

### 1. Basic Printing
* **Goal:** Understand Go syntax, workspace setup, and basic I/O.
* **What it does:** Prints foundational text directly to the console using the standard `fmt` package.

### 2. Concurrency with AdviceSlip API
* **Goal:** Master Go's powerful concurrency model using **goroutines** and **channels**.
* **What it does:** Fetches advice slips simultaneously from the [AdviceSlip API](https://adviceslip.com) without blocking execution, demonstrating efficient asynchronous data fetching.

### 3. Web Server Integration
* **Goal:** Build a functional HTTP server using Go's built-in network packages.
* **What it does:** Spins up a local web server that routes requests. Navigating to the `/advice` endpoint serves a freshly fetched piece of advice directly to the browser.

---

## 🛠️ Tech Stack & Concepts Covered

* **Language:** Go (Golang)
* **Standard Libraries:** `fmt`, `net/http`, `encoding/json`
* **Core Concepts:** Goroutines, Channels, HTTP Routing, API Consumption

---

## 📦 How to Run the Project Local

### Prerequisites
Make sure you have [Go installed](https://go.dev) on your machine.

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com
   cd YOUR_REPO_NAME
   ```

2. Run the web server application:
   ```bash
   go run main.go
   ```
   *(Note: Adjust the file name if your files are split into separate directories like `server/main.go`)*

### Testing the Web Server
Once the server is running, open your browser or use `curl` to visit:
```text
http://localhost:8080/advice
```

---

## 📈 Next Steps
* [ ] Add custom middleware for logging requests.
* [ ] Implement a frontend UI to display the advice cleanly.
* [ ] Add unit tests for the API fetching logic.
