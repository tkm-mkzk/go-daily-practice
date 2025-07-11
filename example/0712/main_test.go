package main

import (
	"reflect"
	"testing"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		expected map[string]int
	}{
		{
			name: "基本",
			text: "Hello world hello Go go programming",
			expected: map[string]int{
				"hello":       2,
				"world":       1,
				"go":          2,
				"programming": 1,
			},
		},
		{
			name:     "空文字列",
			text:     "",
			expected: map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countWords(tt.text)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("countWords(%v) = %v, expected %v", tt.text, result, tt.expected)
			}
		})
	}
}
