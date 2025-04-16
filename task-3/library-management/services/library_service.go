package services

import (
	"fmt"
	"library-management/models"
)

// LibraryManager interface defines the library operations
type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) ([]models.Book, error)
}

// Library implements the LibraryManager interface
type Library struct {
	books   map[int]models.Book   // Map of book ID to Book
	members map[int]models.Member // Map of member ID to Member
}

// NewLibrary creates a new Library instance
func NewLibrary() *Library {
	return &Library{
		books:   make(map[int]models.Book),
		members: make(map[int]models.Member),
	}
}

// AddBook adds a new book to the library
func (l *Library) AddBook(book models.Book) {
	book.Status = "Available"
	l.books[book.ID] = book
}

// RemoveBook removes a book from the library
func (l *Library) RemoveBook(bookID int) error {
	if _, exists := l.books[bookID]; !exists {
		return fmt.Errorf("book with ID %d not found", bookID)
	}
	delete(l.books, bookID)
	return nil
}

// BorrowBook allows a member to borrow a book
func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := l.books[bookID]
	if !exists {
		return fmt.Errorf("book with ID %d not found", bookID)
	}

	if book.Status != "Available" {
		return fmt.Errorf("book with ID %d is already borrowed", bookID)
	}

	member, exists := l.members[memberID]
	if !exists {
		return fmt.Errorf("member with ID %d not found", memberID)
	}

	// Update book status
	book.Status = "Borrowed"
	l.books[bookID] = book

	// Add book to member's borrowed books
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.members[memberID] = member

	return nil
}

// ReturnBook allows a member to return a borrowed book
func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, exists := l.books[bookID]
	if !exists {
		return fmt.Errorf("book with ID %d not found", bookID)
	}

	member, exists := l.members[memberID]
	if !exists {
		return fmt.Errorf("member with ID %d not found", memberID)
	}

	// Check if member has this book borrowed
	found := false
	for i, b := range member.BorrowedBooks {
		if b.ID == bookID {
			// Remove from borrowed books
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("member with ID %d hasn't borrowed book with ID %d", memberID, bookID)
	}

	// Update book status
	book.Status = "Available"
	l.books[bookID] = book
	l.members[memberID] = member

	return nil
}

// ListAvailableBooks returns all available books
func (l *Library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range l.books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

// ListBorrowedBooks returns all books borrowed by a member
func (l *Library) ListBorrowedBooks(memberID int) ([]models.Book, error) {
	member, exists := l.members[memberID]
	if !exists {
		return nil, fmt.Errorf("member with ID %d not found", memberID)
	}
	return member.BorrowedBooks, nil
}
