package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

// THIS SHIT IS TEST DONT PUSH IT !!!!!!!!!!!!!!!!!!
func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

// THIS SHIT IS TEST DONT PUSH IT !!!!!!!!!!!!!!!!!!
func incomingHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w)

	log.Println("Received request:", r.Method)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Println("Processing POST request")

	// Here you can add your actual processing logic
	html := "<p>PDF generated successfully</p>"
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

// THIS SHIT IS TEST DONT PUSH IT !!!!!!!!!!!!!!!!!!
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/pdf", incomingHandler)

	// Enable CORS
	handler := cors.Default().Handler(mux)

	log.Println("Server started on :5501")
	log.Fatal(http.ListenAndServe(":5501", handler))
}

// THIS SHIT IS TEST DONT PUSH IT !!!!!!!!!!!!!!!!!!
