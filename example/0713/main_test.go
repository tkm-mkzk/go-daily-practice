package main

import (
	"reflect"
	"testing"
)

func TestMergeSortedSlices(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []int
		slice2   []int
		expected []int
	}{
		{
			name:     "両方のスライスに要素がある場合",
			slice1:   []int{1, 3, 5, 7},
			slice2:   []int{2, 4, 6, 8},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:     "重複する値がある場合",
			slice1:   []int{1, 3, 5},
			slice2:   []int{3, 6, 9},
			expected: []int{1, 3, 3, 5, 6, 9},
		},
		{
			name:     "slice1が空の場合",
			slice1:   []int{},
			slice2:   []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "slice2が空の場合",
			slice1:   []int{1, 2, 3},
			slice2:   []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "両方のスライスが空の場合",
			slice1:   []int{},
			slice2:   []int{},
			expected: []int{},
		},
		{
			name:     "片方が単一要素の場合",
			slice1:   []int{5},
			slice2:   []int{1, 3, 7, 9},
			expected: []int{1, 3, 5, 7, 9},
		},
		{
			name:     "同じ要素が多数ある場合",
			slice1:   []int{1, 1, 2, 2},
			slice2:   []int{1, 2, 3, 3},
			expected: []int{1, 1, 1, 2, 2, 2, 3, 3},
		},
		{
			name:     "負の数を含む場合",
			slice1:   []int{-3, -1, 1, 3},
			slice2:   []int{-2, 0, 2, 4},
			expected: []int{-3, -2, -1, 0, 1, 2, 3, 4},
		},
		{
			name:     "片方が他方より大きな値のみ含む場合",
			slice1:   []int{1, 2, 3},
			slice2:   []int{10, 20, 30},
			expected: []int{1, 2, 3, 10, 20, 30},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeSortedSlices(tt.slice1, tt.slice2)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MergeSortedSlices(%v, %v) = %v, expected %v",
					tt.slice1, tt.slice2, result, tt.expected)
			}
		})
	}
}

// ベンチマークテスト
func BenchmarkMergeSortedSlices(b *testing.B) {
	slice1 := make([]int, 1000)
	slice2 := make([]int, 1000)

	// テスト用データを準備（偶数は slice1、奇数は slice2）
	for i := 0; i < 1000; i++ {
		slice1[i] = i * 2
		slice2[i] = i*2 + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSortedSlices(slice1, slice2)
	}
}

// 大きなデータでのテスト
func TestMergeSortedSlicesLarge(t *testing.T) {
	slice1 := make([]int, 10000)
	slice2 := make([]int, 10000)

	for i := 0; i < 10000; i++ {
		slice1[i] = i * 2
		slice2[i] = i*2 + 1
	}

	result := MergeSortedSlices(slice1, slice2)

	// 結果の長さをチェック
	if len(result) != 20000 {
		t.Errorf("Expected length 20000, got %d", len(result))
	}

	// ソートされているかチェック
	for i := 1; i < len(result); i++ {
		if result[i] < result[i-1] {
			t.Errorf("Result is not sorted at index %d: %d > %d", i, result[i-1], result[i])
		}
	}
}
