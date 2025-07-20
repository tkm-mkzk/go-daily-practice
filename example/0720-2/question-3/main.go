package main

import "fmt"

func processValue(value any) {
	// ここにコードを追加してください
	// valueがstring型の場合: "文字列: [値]" と表示
	// valueがint型の場合: "数値: [値]" と表示
	// valueがbool型の場合: "真偽値: [値]" と表示
	// その他の型の場合: "不明な型: [値]" と表示

	// 型アサーションを使って実装してください
	if str, ok := value.(string); ok {
		fmt.Printf("文字列: [%v]", str)
		return
	} else if str, ok := value.(int); ok {
		fmt.Printf("数値: [%d]", str)
		return
	} else if str, ok := value.(bool); ok {
		fmt.Printf("真偽値: [%v]", str)
		return
	} else {
		fmt.Printf("不明な型: [%v]", value)
		return
	}
}

func main() {
	processValue("hello")
	processValue(123)
	processValue(true)
	processValue([]int{1, 2, 3})
}
