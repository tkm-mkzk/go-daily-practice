package main

import (
	"errors"
	"fmt"
)

// Calculator は計算機の構造体
type Calculator struct {
	history []string // 計算履歴を保存するスライス
}

// Add は加算を行うメソッド
func (c *Calculator) Add(a, b int) int {
	result := a + b
	// 履歴に追加
	c.history = append(c.history, fmt.Sprintf("%d + %d = %d", a, b, result))
	return result
}

// Subtract は減算を行うメソッド
func (c *Calculator) Subtract(a, b int) int {
	result := a - b
	// 履歴に追加
	c.history = append(c.history, fmt.Sprintf("%d - %d = %d", a, b, result))
	return result
}

// Multiply は乗算を行うメソッド
func (c *Calculator) Multiply(a, b int) int {
	result := a * b
	// 履歴に追加
	c.history = append(c.history, fmt.Sprintf("%d * %d = %d", a, b, result))
	return result
}

// Divide は除算を行うメソッド
func (c *Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("ゼロで割ることはできません")
	}

	result := a / b
	// 履歴に追加
	c.history = append(c.history, fmt.Sprintf("%d / %d = %d", a, b, result))
	return result, nil
}

// History は計算履歴を返すメソッド
func (c *Calculator) History() []string {
	// スライスのコピーを返す（元のスライスを保護）
	historyCopy := make([]string, len(c.history))
	copy(historyCopy, c.history)
	return historyCopy
}

// ClearHistory は履歴をクリアするメソッド（おまけ）
func (c *Calculator) ClearHistory() {
	c.history = []string{}
}

func main() {
	calc := Calculator{} // 初期化（historyは空のスライス）

	fmt.Println(calc.Add(5, 3))       // 8
	fmt.Println(calc.Subtract(10, 4)) // 6
	fmt.Println(calc.Multiply(3, 7))  // 21

	result, err := calc.Divide(15, 3)
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Println(result) // 5
	}

	_, err = calc.Divide(10, 0)
	if err != nil {
		fmt.Println("エラー:", err) // エラー: ゼロで割ることはできません
	}

	fmt.Println("\n=== 計算履歴 ===")
	history := calc.History()
	for _, h := range history {
		fmt.Println(h)
	}

	// おまけ: 履歴をクリアしてテスト
	fmt.Println("\n=== 履歴クリア後 ===")
	calc.ClearHistory()
	calc.Add(100, 200)
	for _, h := range calc.History() {
		fmt.Println(h)
	}
}
