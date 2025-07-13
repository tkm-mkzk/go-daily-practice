package main

import "fmt"

type Book struct {
	Title       string
	Author      string
	Pages       int
	IsAvailable bool
}

type Library struct {
	Books []Book
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book)
}

func (l Library) FindBooksByAuthor(author string) []Book {
	result := []Book{}

	for _, book := range l.Books {
		if book.Author == author {
			result = append(result, book)
		}
	}

	return result
}

func (l *Library) BorrowBook(title string) bool {
	for i, book := range l.Books {
		if book.Title == title && book.IsAvailable {
			l.Books[i].IsAvailable = false
			return true
		}
	}

	return false
}

func main() {
	var Library1 = Library{
		Books: []Book{
			{
				Title:       "Title1",
				Author:      "Author1",
				Pages:       1,
				IsAvailable: true,
			},
			{
				Title:       "Title2",
				Author:      "Author1",
				Pages:       1,
				IsAvailable: false,
			},
			{
				Title:       "Title3",
				Author:      "Author2",
				Pages:       1,
				IsAvailable: true,
			},
			{
				Title:       "Title4",
				Author:      "Author3",
				Pages:       1,
				IsAvailable: true,
			},
			{
				Title:       "Title5",
				Author:      "Author4",
				Pages:       1,
				IsAvailable: false,
			},
		},
	}

	Author1Books := Library1.FindBooksByAuthor("Author1")

	fmt.Println(Author1Books)

	Title4IsAvailable := Library1.BorrowBook("Title4")

	fmt.Println(Title4IsAvailable)
}
