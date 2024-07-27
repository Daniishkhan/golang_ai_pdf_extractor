Here's a README for the project based on the provided code:

# PDF First Name Extractor

This project is a Go-based web service that extracts any given information from PDF documents using Google's Vertex AI and the Gemini model. For demo purposes, it extracts the first name from the PDF.

## Features

- Accepts PDF files via a POST request
- Extracts text from the uploaded PDF
- Uses Google's Vertex AI Gemini model to identify the first name in the extracted text
- Returns the extracted first name in JSON format

## Prerequisites

- Go 1.22.4 or later
- Google Cloud Platform account with Vertex AI API enabled
- Environment variables set up (see Configuration section)

## Installation

1. Clone the repository:

```
git clone git@github.com:Daniishkhan/golang_ai_pdf_extractor.git
cd golang_ai_pdf_extractor
```

2. Install dependencies:

```
go mod download
```

## Configuration

Create a `.env` file in the project root with the following variables:

```
LOCATION=<your-gcp-location>
MODEL_NAME=<gemini-model-name>
PROJECT_ID=<your-gcp-project-id>
```

## Usage

1. Start the server:

```
go run .
```

2. The server will start on port 4000.

3. To parse a PDF, send a POST request to `/parse` with the PDF file in the request body:

```
curl -X POST -F "pdf=@path/to/your/file.pdf" http://localhost:4000/parse
```

4. The server will respond with a JSON object containing the extracted first name:

```json
{
  "first_name": "John"
}
```

## API Endpoints

- `GET /`: Welcome message
- `POST /parse`: Upload and parse a PDF file

## Dependencies

- github.com/go-chi/chi/v5: HTTP router
- cloud.google.com/go/vertexai/genai: Google Vertex AI client
- github.com/joho/godotenv: Environment variable loader
- github.com/ledongthuc/pdf: PDF text extraction

## Error Handling

The service includes error handling for various scenarios, including:
- Invalid file uploads
- PDF parsing errors
- AI model errors

Errors are logged and appropriate HTTP status codes are returned.

## Security Considerations

- Ensure your Google Cloud credentials are kept secure and not committed to version control.
- Consider implementing rate limiting and authentication for production use.
