package main

import (
	"regexp"
	"strings"
)

func translateToJapanese(code string) string {
	lines := strings.Split(code, "\n")
	result := ""
	for _, line := range lines {
		line = strings.TrimSpace(line)
		// 新たに関数定義にも対応
		r := regexp.MustCompile(`func (.+)\((.*)\)`)
		rErr := regexp.MustCompile(`if err != nil {`)
		if strings.HasPrefix(line, "fmt.Println") {
			contents := strings.TrimPrefix(line, "fmt.Println")
			contents = strings.TrimSpace(contents)
			contents = strings.Trim(contents, "\"()")
			result += "「" + contents + "」を出力\n"
		} else if strings.HasPrefix(line, "if") {
			condition := strings.TrimPrefix(line, "if")
			condition = strings.TrimSpace(condition)
			if rErr.MatchString(line) {
				result += "エラーがある場合\n"
			} else {
				result += "もし" + condition + "の場合\n"
			}
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
		} else if r.MatchString(line) {
			matches := r.FindStringSubmatch(line)
			funcName := matches[1]
			params := matches[2]
			result += funcName + "という名前の関数を定義し、パラメータは" + params + "です\n"
		} else {
			result += line + "\n"
		}
	}
	return result
}
