package application

import (
	"log"
	"net/http"
)

func Run(port string) {
	mux := http.NewServeMux()

	mux.Handle("/api/v1/calculate", logResponse(logRequest(http.HandlerFunc(ExpressionHandler))))

	log.Printf("Listening on port %s", port)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
