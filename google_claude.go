package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/vertexai/genai"
	"github.com/joho/godotenv"
	"github.com/ledongthuc/pdf"
)

func extractTextFromPDF(pdfContent []byte) (string, error) {
	reader := bytes.NewReader(pdfContent)
	pdfReader, err := pdf.NewReader(reader, int64(len(pdfContent)))
	if err != nil {
		return "", fmt.Errorf("error creating PDF reader: %w", err)
	}

	var buf bytes.Buffer
	b, err := pdfReader.GetPlainText()
	if err != nil {
		return "", fmt.Errorf("error extracting text from PDF: %w", err)
	}

	buf.ReadFrom(b)
	return buf.String(), nil
}

func parsePDF(pdfContent []byte, fileName string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	location := os.Getenv("LOCATION")
	modelName := os.Getenv("MODEL_NAME")
	projectID := os.Getenv("PROJECT_ID")

	if location == "" || modelName == "" || projectID == "" {
		return "", fmt.Errorf("missing required environment variables")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, projectID, location)
	if err != nil {
		return "", fmt.Errorf("error creating client: %w", err)
	}
	defer client.Close()

	extractedText, err := extractTextFromPDF(pdfContent)
	if err != nil {
		return "", fmt.Errorf("error extracting text from PDF: %w", err)
	}
	log.Printf("Extracted text (first 200 chars): %s", extractedText[:min(200, len(extractedText))])

	gemini := client.GenerativeModel(modelName)
	prompt := genai.Text(fmt.Sprintf("Extract the first name from the following text extracted from a PDF. Respond with only the first name in JSON format like this: {\"firstName\": \"John\"}. If you can't find a name, respond with {\"firstName\": \"\"} and nothing else. Here is the extracted text: %s", extractedText))

	resp, err := gemini.GenerateContent(ctx, prompt)
	if err != nil {
		return "", fmt.Errorf("error generating content: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("no content generated")
	}

	firstName := resp.Candidates[0].Content.Parts[0].(genai.Text)
	log.Printf("Gemini response: %s", firstName)
	return string(firstName), nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
