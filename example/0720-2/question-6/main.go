// 課題：
// 1. nil の *Document を Printer インターフェースに代入
// 2. インターフェース変数が nil かどうかをチェック
// 3. 結果を確認して、なぜそうなるのかを考えてください
// この問題で「インターフェースとnil」の微妙な関係を学びましょう！

package main

import "fmt"

type Printer interface {
	Print()
}

type Document struct {
	content string
}

func (d *Document) Print() {
	if d == nil {
		fmt.Println("空のドキュメントです")
		return
	}
	fmt.Println("内容:", d.content)
}

func main() {
	var doc *Document = nil

	// 1. docをPrinter型の変数に代入してください
	var printer Printer = doc

	// 2. そのPrinter変数がnilかどうかチェックしてください
	var isPrinterIsNil bool
	if v, ok := printer.(*Document); !ok || v == nil {
		isPrinterIsNil = true
	} else {
		isPrinterIsNil = false
	}
	fmt.Println(isPrinterIsNil)

	// 3. Printメソッドを呼び出してください
	doc.Print()

	// 期待する動作を確認してみてください
	doc = &Document{}
	doc.content = "テスト"
	doc.Print()
}
