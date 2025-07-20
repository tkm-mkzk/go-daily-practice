package main

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

// ReadWriterインターフェースを定義してください
// ReaderとWriterの両方を埋め込んで作成してください

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	name    string
	content string
}

// FileがReadWriterインターフェースを実装するよう
// 必要なメソッドを追加してください

func (f File) Read() string {
	return f.name + f.content
}

func (f *File) Write(data string) {
	f.content += data
}

func main() {
	f := File{name: "test.txt", content: ""}

	var rw ReadWriter = &f
	rw.Write("Hello, World!")
	fmt.Println(rw.Read())
}
