package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	reversedString := reverseString(strings.ToLower(s))
	return strings.ToLower(s) == reversedString
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	string := "テストステ"

	result := isPalindrome(string)
	fmt.Println(result)
}
