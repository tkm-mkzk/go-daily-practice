package main

import (
	"errors"
	"testing"
)

func TestBasicCalculatorAdd(t *testing.T) {
	calc := BasicCalculator{}

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"正の数の足し算", 10.0, 5.0, 15.0},
		{"負の数を含む足し算", -5.0, 3.0, -2.0},
		{"小数点の足し算", 2.5, 3.7, 6.2},
		{"ゼロとの足し算", 10.0, 0.0, 10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%.2f, %.2f) = %.2f; expected %.2f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestBasicCalculatorSubtract(t *testing.T) {
	calc := BasicCalculator{}

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"正の数の引き算", 10.0, 3.0, 7.0},
		{"負の数を含む引き算", 5.0, -3.0, 8.0},
		{"小数点の引き算", 7.5, 2.3, 5.2},
		{"同じ数の引き算", 5.0, 5.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Subtract(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Subtract(%.2f, %.2f) = %.2f; expected %.2f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestBasicCalculatorMultiply(t *testing.T) {
	calc := BasicCalculator{}

	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"正の数の掛け算", 4.0, 5.0, 20.0},
		{"負の数を含む掛け算", -3.0, 4.0, -12.0},
		{"小数点の掛け算", 2.5, 4.0, 10.0},
		{"ゼロとの掛け算", 7.0, 0.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calc.Multiply(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Multiply(%.2f, %.2f) = %.2f; expected %.2f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestBasicCalculatorDivide(t *testing.T) {
	calc := BasicCalculator{}

	// 正常ケースのテスト
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"正の数の割り算", 10.0, 2.0, 5.0},
		{"小数点の割り算", 7.5, 2.5, 3.0},
		{"負の数を含む割り算", -12.0, 3.0, -4.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.Divide(tt.a, tt.b)
			if err != nil {
				t.Errorf("Divide(%.2f, %.2f) returned unexpected error: %v", tt.a, tt.b, err)
			}
			if result != tt.expected {
				t.Errorf("Divide(%.2f, %.2f) = %.2f; expected %.2f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestBasicCalculatorDivideByZero(t *testing.T) {
	calc := BasicCalculator{}

	// ゼロ除算のテスト
	tests := []struct {
		name string
		a    float64
	}{
		{"正の数をゼロで割る", 10.0},
		{"負の数をゼロで割る", -5.0},
		{"ゼロをゼロで割る", 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := calc.Divide(tt.a, 0.0)

			// エラーが返されることを確認
			if err == nil {
				t.Errorf("Divide(%.2f, 0.0) should return error, but got result: %.2f", tt.a, result)
			}

			// DivisionByZeroErrorが返されることを確認
			var divErr DivisionByZeroError
			if !errors.As(err, &divErr) {
				t.Errorf("Divide(%.2f, 0.0) should return DivisionByZeroError, but got: %T", tt.a, err)
			}

			// エラーメッセージの確認
			expectedMsg := "division by zero is not allowed"
			if err.Error() != expectedMsg {
				t.Errorf("Error message = '%s'; expected '%s'", err.Error(), expectedMsg)
			}

			// 結果が0であることを確認
			if result != 0 {
				t.Errorf("Divide(%.2f, 0.0) should return 0 as result, but got: %.2f", tt.a, result)
			}
		})
	}
}

func TestDivisionByZeroError(t *testing.T) {
	err := DivisionByZeroError{}

	// Error()メソッドのテスト
	expectedMsg := "division by zero is not allowed"
	if err.Error() != expectedMsg {
		t.Errorf("DivisionByZeroError.Error() = '%s'; expected '%s'", err.Error(), expectedMsg)
	}

	// error インターフェースを満たすことを確認
	var e error = err
	if e.Error() != expectedMsg {
		t.Errorf("DivisionByZeroError as error interface = '%s'; expected '%s'", e.Error(), expectedMsg)
	}
}

func TestCalculatorInterface(t *testing.T) {
	// BasicCalculatorがCalculatorインターフェースを満たすことをテスト
	var calc Calculator = BasicCalculator{}

	// インターフェースを通じてメソッドを呼び出し
	addResult := calc.Add(5.0, 3.0)
	if addResult != 8.0 {
		t.Errorf("Calculator interface Add(5.0, 3.0) = %.2f; expected 8.00", addResult)
	}

	subResult := calc.Subtract(10.0, 4.0)
	if subResult != 6.0 {
		t.Errorf("Calculator interface Subtract(10.0, 4.0) = %.2f; expected 6.00", subResult)
	}

	mulResult := calc.Multiply(3.0, 4.0)
	if mulResult != 12.0 {
		t.Errorf("Calculator interface Multiply(3.0, 4.0) = %.2f; expected 12.00", mulResult)
	}

	divResult, err := calc.Divide(15.0, 3.0)
	if err != nil {
		t.Errorf("Calculator interface Divide(15.0, 3.0) returned unexpected error: %v", err)
	}
	if divResult != 5.0 {
		t.Errorf("Calculator interface Divide(15.0, 3.0) = %.2f; expected 5.00", divResult)
	}
}

func TestPerformCalculations(t *testing.T) {
	// PerformCalculations関数のテスト（出力の確認は省略、エラーが起きないことを確認）
	calc := BasicCalculator{}

	// 正常ケース（panic などが起きないことを確認）
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("PerformCalculations panicked with normal inputs: %v", r)
		}
	}()

	PerformCalculations(calc, 10.0, 2.0)
	PerformCalculations(calc, 10.0, 0.0) // ゼロ除算ケース
	PerformCalculations(calc, -5.0, 3.0) // 負の数ケース
}
