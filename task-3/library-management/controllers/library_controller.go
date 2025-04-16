package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"library-management/models"
	"library-management/services"
)

type LibraryController struct {
	library services.LibraryManager
}

func NewLibraryController(library services.LibraryManager) *LibraryController {
	return &LibraryController{library: library}
}

func (lc *LibraryController) Run() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books by Member")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")

		scanner.Scan()
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			lc.addBook(scanner)
		case 2:
			lc.removeBook(scanner)
		case 3:
			lc.borrowBook(scanner)
		case 4:
			lc.returnBook(scanner)
		case 5:
			lc.listAvailableBooks()
		case 6:
			lc.listBorrowedBooks(scanner)
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		fmt.Print("\nDo you want to continue? Press any key and hit Enter...")
		scanner.Scan()
	}
}

func (lc *LibraryController) addBook(scanner *bufio.Scanner) {
	fmt.Print("Enter Book ID: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Enter Book Title: ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Print("Enter Book Author: ")
	scanner.Scan()
	author := scanner.Text()

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
	}

	lc.library.AddBook(book)
	fmt.Println("Book added successfully!")
}

func (lc *LibraryController) removeBook(scanner *bufio.Scanner) {
	fmt.Print("Enter Book ID to remove: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	err := lc.library.RemoveBook(id)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book removed successfully!")
	}
}

func (lc *LibraryController) borrowBook(scanner *bufio.Scanner) {
	fmt.Print("Enter Book ID to borrow: ")
	scanner.Scan()
	bookID, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Enter Member ID: ")
	scanner.Scan()
	memberID, _ := strconv.Atoi(scanner.Text())

	err := lc.library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed successfully!")
	}
}

func (lc *LibraryController) returnBook(scanner *bufio.Scanner) {
	fmt.Print("Enter Book ID to return: ")
	scanner.Scan()
	bookID, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Enter Member ID: ")
	scanner.Scan()
	memberID, _ := strconv.Atoi(scanner.Text())

	err := lc.library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned successfully!")
	}
}

func (lc *LibraryController) listAvailableBooks() {
	books := lc.library.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No available books.")
		return
	}

	fmt.Println("\n+------+------------------------------+--------------------------+")
	fmt.Println("|  ID  |           Title              |         Author          |")
	fmt.Println("+------+------------------------------+--------------------------+")

	for _, book := range books {
		fmt.Printf("| %-4d | %-28s | %-24s |\n", book.ID, book.Title, book.Author)
	}

	fmt.Println("+------+------------------------------+--------------------------+")
}

func (lc *LibraryController) listBorrowedBooks(scanner *bufio.Scanner) {
	fmt.Print("Enter Member ID: ")
	scanner.Scan()
	memberID, _ := strconv.Atoi(scanner.Text())

	books, err := lc.library.ListBorrowedBooks(memberID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("No books borrowed by this member.")
		return
	}

	fmt.Println("\n+------+------------------------------+--------------------------+")
	fmt.Println("|  ID  |           Title              |         Author          |")
	fmt.Println("+------+------------------------------+--------------------------+")

	for _, book := range books {
		fmt.Printf("| %-4d | %-28s | %-24s |\n", book.ID, book.Title, book.Author)
	}

	fmt.Println("+------+------------------------------+--------------------------+")
}