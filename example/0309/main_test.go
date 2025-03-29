package main

import (
	"testing"
)

func TestPersonString(t *testing.T) {
	// テストケース1: 基本的なケース
	p := Person{Name: "太郎", Age: 25}
	expected := "太郎-----25"
	actual := p.String()

	if actual != expected {
		t.Errorf("Person.String() = %q, want %q", actual, expected)
	}
}

func TestPersonStringWithDifferentValues(t *testing.T) {
	// テストケース2: 異なる値でのテスト
	testCases := []struct {
		name     string
		person   Person
		expected string
	}{
		{
			name:     "花子のテスト",
			person:   Person{Name: "花子", Age: 30},
			expected: "花子-----30",
		},
		{
			name:     "空の名前",
			person:   Person{Name: "", Age: 0},
			expected: "-----0",
		},
		{
			name:     "長い名前",
			person:   Person{Name: "田中一郎", Age: 100},
			expected: "田中一郎-----100",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.person.String()
			if actual != tc.expected {
				t.Errorf("Person.String() = %q, want %q", actual, tc.expected)
			}
		})
	}
}

func TestStringerInterface(t *testing.T) {
	// テストケース3: Stringerインターフェースとしてのテスト
	p := Person{Name: "次郎", Age: 35}
	var stringer Stringer = p

	expected := "次郎-----35"
	actual := stringer.String()

	if actual != expected {
		t.Errorf("Stringer.String() = %q, want %q", actual, expected)
	}
}

func TestPersonFields(t *testing.T) {
	// テストケース4: Person構造体のフィールドテスト
	p := Person{Name: "三郎", Age: 40}

	if p.Name != "三郎" {
		t.Errorf("Person.Name = %q, want %q", p.Name, "三郎")
	}

	if p.Age != 40 {
		t.Errorf("Person.Age = %d, want %d", p.Age, 40)
	}
}
