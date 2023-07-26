package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
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

func translateToJapanese(code string) string {
	lines := strings.Split(code, "\n")
	result := ""
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "fmt.Println") {
			contents := strings.TrimPrefix(line, "fmt.Println")
			contents = strings.TrimSpace(contents)
			contents = strings.Trim(contents, "\"()")
			result += "「" + contents + "」を出力\n"
		} else if strings.HasPrefix(line, "if") {
			condition := strings.TrimPrefix(line, "if")
			condition = strings.TrimSpace(condition)
			result += "もし" + condition + "の場合\n"
		} else if strings.HasPrefix(line, "for") {
			condition := strings.TrimPrefix(line, "for")
			condition = strings.TrimSpace(condition)
			result += condition + "の間、以下の処理を繰り返す\n"
		} else if strings.HasPrefix(line, "import") {
			packages := strings.TrimPrefix(line, "import")
			packages = strings.TrimSpace(packages)
			packages = strings.Trim(packages, "\"()")
			result += packages + "パッケージをインポート\n"
		} else if strings.Contains(line, ":=") {
			parts := strings.Split(line, ":=")
			variable := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			if strings.Contains(value, "[]int") {
				result += variable + "という名前の整数型スライスを定義し、値は" + strings.Trim(value, "[]int{}") + "です\n"
			} else {
				result += variable + "という名前の変数を定義し、値は" + value + "です\n"
			}
		} else {
			result += line + "\n"
		}
	}
	return result
}

func main() {
	http.HandleFunc("/translate", handleTranslateRequest)

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
