package main

import (
	"fmt"
	"net/http"

	"github.com/carsonkiibi/pdfapp/backend/process"
	"github.com/carsonkiibi/pdfapp/backend/process/commands"
)

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

func main() {
	str := "I like {B,S14:apples} {U, I:bananas}"
	outTokens := process.ProcessInput(str)
	http.HandleFunc("/pdf", pdfHandler(outTokens))
	fmt.Println("Serving on port 8080")
	http.ListenAndServe(":8080", nil)
}
