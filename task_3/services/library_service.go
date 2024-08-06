package services

import (
	"errors"

	"github.com/Ephrem-shimels21/models"
)

type LibraryManager interface {
	AddBook(book models.Book) error
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() ([]models.Book, error)
	ListBorrowedBooks(memberID int) ([]models.Book, error)
}

type Library struct {
	books   map[int]models.Book
	Members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

func (lib *Library) AddMember(member models.Member) error {
	if _, ok := lib.Members[member.ID]; ok {
		return errors.New("member already exists")
	}
	lib.Members[member.ID] = member
	return nil

}

func (lib *Library) AddBook(book models.Book) error {
	if _, ok := lib.books[book.ID]; ok {
		return errors.New("book already exists")
	}
	lib.books[book.ID] = book
	return nil
}

func (lib *Library) RemoveBook(bookID int) error {
	if _, ok := lib.books[bookID]; !ok {
		return errors.New("book not found")
	}
	delete(lib.books, bookID)
	return nil
}

func (lib *Library) BorrowBook(bookID int, memberID int) error {
	book, ok := lib.books[bookID]

	if !ok {
		return errors.New("book not found")
	}

	if book.Status == "borrowed" {
		return errors.New("book is already borrowed")

	}

	member, ok := lib.Members[memberID]

	if !ok {
		return errors.New("member Not found")
	}

	book.Status = "borrowed"
	lib.books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	lib.Members[memberID] = member
	return nil

}

func (lib *Library) ReturnBook(bookID int, memberID int) error {
	book, ok := lib.books[bookID]

	if !ok || book.Status == "available" {
		return errors.New("book is not borrowed")
	}

	member, ok := lib.Members[memberID]

	if !ok {
		return errors.New("member Not found")
	}

	for i, borrowedBook := range member.BorrowedBooks {
		if borrowedBook.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			book.Status = "Available"
			lib.books[bookID] = book
			lib.Members[memberID] = member
			return nil
		}
	}
	return errors.New("member has not borrowed this book")

}

func (l *Library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range l.books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := l.Members[memberID]
	if !exists {
		return nil
	}
	return member.BorrowedBooks
}
