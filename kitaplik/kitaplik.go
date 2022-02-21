package kitaplik

import "strings"

func Search(texts []string) []string {
	fullText := strings.Join(texts, " ")
	lowerFullText := strings.ToLower(fullText)
	result := []string{}
	for _, book := range books {
		if strings.Contains(strings.ToLower(book), lowerFullText) {
			result = append(result, book)
		}
	}
	return result
}

func List() []string {
	return books
}
