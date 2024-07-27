package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ParseRequest struct {
	PDFContent []byte `json:"pdf_content"`
	FileName   string `json:"file_name"`
}

type ParseResponse struct {
	FirstName string `json:"first_name"`
}

func Routes(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Post("/parse", func(w http.ResponseWriter, r *http.Request) {
		// Parse the multipart form data
		err := r.ParseMultipartForm(32 << 20) // 10 MB max
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to parse form: %v", err), http.StatusBadRequest)
			return
		}

		// Get the file from the form data
		file, header, err := r.FormFile("pdf")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error retrieving the file: %v", err), http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Log file details
		fmt.Printf("Uploaded File: %+v\n", header.Filename)
		fmt.Printf("File Size: %+v\n", header.Size)
		fmt.Printf("MIME Header: %+v\n", header.Header)

		pdfContent, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading the file: %v", err), http.StatusBadRequest)
			return
		}

		response, err := parsePDF(pdfContent, header.Filename)
		if err != nil {
			log.Printf("Error parsing PDF: %v", err)
			http.Error(w, "Error parsing PDF", http.StatusInternalServerError)
			return
		}

		var result map[string]string
		err = json.Unmarshal([]byte(response), &result)
		if err != nil {
			log.Printf("Error unmarshaling response: %v", err)
			http.Error(w, "Error processing response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ParseResponse{FirstName: result["firstName"]})
	})
}
