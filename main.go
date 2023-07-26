package main

import (
	"fmt"
	"io"
	"net/http"
)

func handleTranslateRequest(w http.ResponseWriter, r *http.Request) {
	// Read the program text from the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	programText := string(body)

	// Translate the program text
	description := translateToJapanese(programText)

	// Write the description as a response
	fmt.Fprint(w, description)
}

func main() {
	http.HandleFunc("/translate", handleTranslateRequest)

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
