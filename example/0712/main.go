package main

import (
	"fmt"
	"strings"
)

// func countWords(text string) map[string]int {
// 	result := make(map[string]int)
// 	segmaentedTexts := strings.Fields(strings.ToLower(text))
// 	texts := make([]string, 0, len(segmaentedTexts))

// 	for _, text := range segmaentedTexts {
// 		if !slices.Contains(texts, text) {
// 			texts = append(texts, text)
// 			result[text] = 1
// 		} else if slices.Contains(texts, text) {
// 			result[text] = +1
// 		}
// 	}

// 	return result
// }

// func countWords(text string) map[string]int {
// 	result := make(map[string]int)
// 	segmaentedTexts := strings.Fields(strings.ToLower(text))

// 	for _, word := range segmaentedTexts {
// 		result[word]++
// 	}

//		return result
//	}

func countWords(text string) map[string]int {
	result := make(map[string]int)
	segmaentedTexts := strings.Fields(strings.ToLower(text))

	for _, word := range segmaentedTexts {
		if count, exists := result[word]; exists {
			result[word] = count + 1
		} else {
			result[word] = 1
		}
	}

	return result
}

func main() {
	text := "Hello world hello Go go programming"
	result := countWords(text)

	fmt.Printf("入力: %s\n", text)
	fmt.Printf("結果: %v\n", result)

	// より見やすく表示
	fmt.Println("詳細:")
	for word, count := range result {
		fmt.Printf("  %s: %d回\n", word, count)
	}
}
