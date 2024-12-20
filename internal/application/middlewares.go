package application

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL)
		log.Printf("Headers: %v", r.Header)

		if r.Method == http.MethodPost {
			var body map[string]interface{}

			bodyBytes, _ := io.ReadAll(r.Body)

			if err := json.Unmarshal(bodyBytes, &body); err == nil {
				log.Printf("Request Body: %v", body)
			}

			r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		}

		next.ServeHTTP(w, r)
	})
}

func logResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &ResponseRecorder{ResponseWriter: w, StatusCode: http.StatusOK}
		next.ServeHTTP(rec, r)

		log.Printf("Response: %d", rec.StatusCode)
	})
}
