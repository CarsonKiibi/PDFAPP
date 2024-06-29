package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/carsonkiibi/pdfapp/backend/process"
	"github.com/carsonkiibi/pdfapp/backend/process/commands"
	"github.com/rs/cors"
)

type TextRequest struct {
	Text string `json:"text"`
}

type PDFResponse struct {
	PDFData []byte `json:"pdf_data"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func pdfHandler(tokens []commands.Token) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pdfBytes, err := process.GeneratePDF(tokens)
		if err != nil {
			http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
			return
		}
		enableCors(&w)
		w.Header().Set("Content-Type", "application/pdf")
		w.Header().Set("Content-Disposition", "inline; filename=document.pdf")
		w.Write(pdfBytes)
	}

}

func incomingHandler(w http.ResponseWriter, r *http.Request) {
	var req TextRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tokens := process.ProcessInput(req.Text)
	pdfBytes, err := process.GeneratePDF(tokens)
	if err != nil {
		http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
		return
	}
	jsonBytes := PDFResponse{PDFData: pdfBytes}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonBytes)
}

func main() {
	// str := "I like {B,S14:apples}"
	// outTokens := process.ProcessInput(str)
	mux := http.NewServeMux()
	mux.HandleFunc("/pdf", incomingHandler)

	// Enable CORS
	handler := cors.Default().Handler(mux)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
