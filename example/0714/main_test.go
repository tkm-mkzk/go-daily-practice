package main

import (
	"testing"
)

func TestAddBook(t *testing.T) {
	library := Library{}

	book := Book{
		Title:       "Test Book",
		Author:      "Test Author",
		Pages:       100,
		IsAvailable: true,
	}

	// 本を追加前の状態確認
	if len(library.Books) != 0 {
		t.Errorf("初期状態で本の数が0でない: got %d", len(library.Books))
	}

	// 本を追加
	library.AddBook(book)

	// 追加後の確認
	if len(library.Books) != 1 {
		t.Errorf("本の追加後に本の数が1でない: got %d", len(library.Books))
	}

	if library.Books[0].Title != "Test Book" {
		t.Errorf("追加した本のタイトルが正しくない: got '%s', expected 'Test Book'", library.Books[0].Title)
	}

	if library.Books[0].Author != "Test Author" {
		t.Errorf("追加した本の著者が正しくない: got '%s', expected 'Test Author'", library.Books[0].Author)
	}

	if library.Books[0].Pages != 100 {
		t.Errorf("追加した本のページ数が正しくない: got %d, expected 100", library.Books[0].Pages)
	}

	if !library.Books[0].IsAvailable {
		t.Error("追加した本が利用可能でない")
	}
}

func TestAddMultipleBooks(t *testing.T) {
	library := Library{}

	books := []Book{
		{Title: "Book1", Author: "Author1", Pages: 100, IsAvailable: true},
		{Title: "Book2", Author: "Author2", Pages: 200, IsAvailable: false},
		{Title: "Book3", Author: "Author1", Pages: 150, IsAvailable: true},
	}

	// 複数の本を追加
	for _, book := range books {
		library.AddBook(book)
	}

	// 本の数を確認
	if len(library.Books) != 3 {
		t.Errorf("3冊追加後の本の数が正しくない: got %d, expected 3", len(library.Books))
	}

	// 各本が正しく追加されているか確認
	for i, expectedBook := range books {
		if library.Books[i].Title != expectedBook.Title {
			t.Errorf("Book[%d] title: got '%s', expected '%s'", i, library.Books[i].Title, expectedBook.Title)
		}
		if library.Books[i].Author != expectedBook.Author {
			t.Errorf("Book[%d] author: got '%s', expected '%s'", i, library.Books[i].Author, expectedBook.Author)
		}
	}
}

func TestFindBooksByAuthor(t *testing.T) {
	library := Library{
		Books: []Book{
			{Title: "Book1", Author: "Author1", Pages: 100, IsAvailable: true},
			{Title: "Book2", Author: "Author1", Pages: 200, IsAvailable: false},
			{Title: "Book3", Author: "Author2", Pages: 150, IsAvailable: true},
			{Title: "Book4", Author: "Author3", Pages: 250, IsAvailable: true},
			{Title: "Book5", Author: "Author1", Pages: 180, IsAvailable: true},
		},
	}

	// Author1の本を検索（3冊あるはず）
	author1Books := library.FindBooksByAuthor("Author1")

	if len(author1Books) != 3 {
		t.Errorf("Author1の本の数が正しくない: got %d, expected 3", len(author1Books))
	}

	// 見つかった本のタイトルを確認
	expectedTitles := []string{"Book1", "Book2", "Book5"}
	for i, book := range author1Books {
		if book.Title != expectedTitles[i] {
			t.Errorf("Author1 book[%d] title: got '%s', expected '%s'", i, book.Title, expectedTitles[i])
		}
		if book.Author != "Author1" {
			t.Errorf("Author1 book[%d] author: got '%s', expected 'Author1'", i, book.Author)
		}
	}

	// Author2の本を検索（1冊あるはず）
	author2Books := library.FindBooksByAuthor("Author2")

	if len(author2Books) != 1 {
		t.Errorf("Author2の本の数が正しくない: got %d, expected 1", len(author2Books))
	}

	if author2Books[0].Title != "Book3" {
		t.Errorf("Author2の本のタイトルが正しくない: got '%s', expected 'Book3'", author2Books[0].Title)
	}
}

func TestFindBooksByAuthorNotFound(t *testing.T) {
	library := Library{
		Books: []Book{
			{Title: "Book1", Author: "Author1", Pages: 100, IsAvailable: true},
		},
	}

	// 存在しない著者で検索
	books := library.FindBooksByAuthor("NonExistentAuthor")

	if len(books) != 0 {
		t.Errorf("存在しない著者の検索結果が空でない: got %d books", len(books))
	}

	// 空のライブラリで検索
	emptyLibrary := Library{}
	books = emptyLibrary.FindBooksByAuthor("Author1")

	if len(books) != 0 {
		t.Errorf("空のライブラリでの検索結果が空でない: got %d books", len(books))
	}
}

func TestBorrowBook(t *testing.T) {
	library := Library{
		Books: []Book{
			{Title: "Available Book", Author: "Author1", Pages: 100, IsAvailable: true},
			{Title: "Unavailable Book", Author: "Author2", Pages: 200, IsAvailable: false},
			{Title: "Another Available Book", Author: "Author3", Pages: 150, IsAvailable: true},
		},
	}

	// 利用可能な本を借りる
	success := library.BorrowBook("Available Book")

	if !success {
		t.Error("利用可能な本を借りることができなかった")
	}

	// 本の状態が変更されているか確認
	if library.Books[0].IsAvailable {
		t.Error("借りた本がまだ利用可能になっている")
	}

	// 他の本の状態が変わっていないことを確認
	if library.Books[1].IsAvailable {
		t.Error("借りていない本(元々unavailable)の状態が変わった")
	}

	if !library.Books[2].IsAvailable {
		t.Error("借りていない本(available)の状態が変わった")
	}
}

func TestBorrowBookUnavailable(t *testing.T) {
	library := Library{
		Books: []Book{
			{Title: "Unavailable Book", Author: "Author1", Pages: 100, IsAvailable: false},
		},
	}

	// 利用不可能な本を借りようとする
	success := library.BorrowBook("Unavailable Book")

	if success {
		t.Error("利用不可能な本を借りることができてしまった")
	}

	// 本の状態が変わっていないことを確認
	if library.Books[0].IsAvailable {
		t.Error("借りられなかった本の状態が変わってしまった")
	}
}

func TestBorrowBookNotFound(t *testing.T) {
	library := Library{
		Books: []Book{
			{Title: "Existing Book", Author: "Author1", Pages: 100, IsAvailable: true},
		},
	}

	// 存在しない本を借りようとする
	success := library.BorrowBook("Non-existent Book")

	if success {
		t.Error("存在しない本を借りることができてしまった")
	}

	// 既存の本の状態が変わっていないことを確認
	if !library.Books[0].IsAvailable {
		t.Error("存在しない本を借りようとして、既存の本の状態が変わってしまった")
	}
}

func TestBorrowBookTwice(t *testing.T) {
	library := Library{
		Books: []Book{
			{Title: "Test Book", Author: "Author1", Pages: 100, IsAvailable: true},
		},
	}

	// 最初に借りる
	success1 := library.BorrowBook("Test Book")
	if !success1 {
		t.Error("最初の借用が失敗した")
	}

	// 同じ本をもう一度借りようとする
	success2 := library.BorrowBook("Test Book")
	if success2 {
		t.Error("既に借りた本をもう一度借りることができてしまった")
	}

	// 本の状態確認
	if library.Books[0].IsAvailable {
		t.Error("二度目の借用試行後、本が利用可能になってしまった")
	}
}

func TestLibraryIntegration(t *testing.T) {
	// 統合テスト：全機能を組み合わせてテスト
	library := Library{}

	// 本を追加
	books := []Book{
		{Title: "Go Programming", Author: "Rob Pike", Pages: 300, IsAvailable: true},
		{Title: "Effective Go", Author: "Rob Pike", Pages: 250, IsAvailable: true},
		{Title: "Clean Code", Author: "Robert Martin", Pages: 400, IsAvailable: true},
		{Title: "Design Patterns", Author: "Gang of Four", Pages: 500, IsAvailable: false},
	}

	for _, book := range books {
		library.AddBook(book)
	}

	// 著者で検索
	robPikeBooks := library.FindBooksByAuthor("Rob Pike")
	if len(robPikeBooks) != 2 {
		t.Errorf("Rob Pikeの本の数が正しくない: got %d, expected 2", len(robPikeBooks))
	}

	// 本を借りる
	success := library.BorrowBook("Go Programming")
	if !success {
		t.Error("Go Programmingを借りることができなかった")
	}

	// 借りた本がもう一度借りられないことを確認
	success = library.BorrowBook("Go Programming")
	if success {
		t.Error("既に借りたGo Programmingをもう一度借りることができてしまった")
	}

	// 利用不可能な本を借りようとする
	success = library.BorrowBook("Design Patterns")
	if success {
		t.Error("利用不可能なDesign Patternsを借りることができてしまった")
	}

	// 別の利用可能な本を借りる
	success = library.BorrowBook("Clean Code")
	if !success {
		t.Error("Clean Codeを借りることができなかった")
	}

	// 最終的な状態確認
	if len(library.Books) != 4 {
		t.Errorf("最終的な本の数が正しくない: got %d, expected 4", len(library.Books))
	}
}
