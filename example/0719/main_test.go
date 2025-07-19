package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"日本語回文", "テストステ", true},
		{"英語回文小文字", "racecar", true},
		{"英語回文大文字混じり", "Racecar", true},
		{"大文字小文字混在", "Madam", true},
		{"回文ではない", "hello", false},
		{"単一文字", "a", true},
		{"空文字", "", true},
		{"2文字回文", "aa", true},
		{"2文字非回文", "ab", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isPalindrome(test.input)
			if result != test.expected {
				t.Errorf("isPalindrome(%q) = %v, expected %v", test.input, result, test.expected)
			}
		})
	}
}
