package main

import "fmt"

type Calculator interface {
	Add(a, b float64) float64
	Subtract(a, b float64) float64
	Multiply(a, b float64) float64
	Divide(a, b float64) (float64, error)
}

type DivisionByZeroError struct{}

func (e DivisionByZeroError) Error() string {
	return "division by zero is not allowed"
}

type BasicCalculator struct{}

func (bc BasicCalculator) Add(a, b float64) float64 {
	return a + b
}

func (bc BasicCalculator) Subtract(a, b float64) float64 {
	return a - b
}

func (bc BasicCalculator) Multiply(a, b float64) float64 {
	return a * b
}

func (bc BasicCalculator) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, DivisionByZeroError{}
	}

	return a / b, nil
}

func PerformCalculations(calc Calculator, a, b float64) {
	fmt.Printf("\n=== 計算結果 (a=%.2f, b=%.2f) ===\n", a, b)

	// 足し算
	result := calc.Add(a, b)
	fmt.Printf("Add: %.2f + %.2f = %.2f\n", a, b, result)

	// 引き算
	result = calc.Subtract(a, b)
	fmt.Printf("Subtract: %.2f - %.2f = %.2f\n", a, b, result)

	// 掛け算
	result = calc.Multiply(a, b)
	fmt.Printf("Multiply: %.2f * %.2f = %.2f\n", a, b, result)

	// 割り算（エラーハンドリング付き）
	result, err := calc.Divide(a, b)
	if err != nil {
		fmt.Printf("Divide: %.2f / %.2f = Error: %v\n", a, b, err)
	} else {
		fmt.Printf("Divide: %.2f / %.2f = %.2f\n", a, b, result)
	}
}

func main() {
	// BasicCalculatorのインスタンスを作成
	calc := BasicCalculator{}

	// 正常なケースをテスト
	PerformCalculations(calc, 10, 2)

	// 0除算のケースをテスト
	PerformCalculations(calc, 10, 0)

	// 追加のテストケース
	PerformCalculations(calc, 15.5, 3.2)
}
