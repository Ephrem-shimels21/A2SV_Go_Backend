package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/Ephrem-shimels21/models"
	"github.com/Ephrem-shimels21/services"
)

func RunConsoleApp(library *services.Library) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Library Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("0. Exit")
		fmt.Print("Choose an option: ")
		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			addBook(scanner, library)
		case "2":
			removeBook(scanner, library)
		case "3":
			borrowBook(scanner, library)
		case "4":
			returnBook(scanner, library)
		case "5":
			listAvailableBooks(library)
		case "6":
			listBorrowedBooks(scanner, library)
		case "0":
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}

func addBook(scanner *bufio.Scanner, library *services.Library) {
	fmt.Print("Enter book ID: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Enter book title: ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Print("Enter book author: ")
	scanner.Scan()
	author := scanner.Text()

	book := models.Book{
		ID:     id,
		Title:  title,
		Author: author,
		Status: "Available",
	}
	library.AddBook(book)
	fmt.Println("Book added.")
}

func removeBook(scanner *bufio.Scanner, library *services.Library) {
	fmt.Print("Enter book ID to remove: ")
	scanner.Scan()
	id, _ := strconv.Atoi(scanner.Text())

	library.RemoveBook(id)
	fmt.Println("Book removed.")
}

func borrowBook(scanner *bufio.Scanner, library *services.Library) {
	fmt.Print("Enter book ID to borrow: ")
	scanner.Scan()
	bookID, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Enter member ID: ")
	scanner.Scan()
	memberID, _ := strconv.Atoi(scanner.Text())

	err := library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed.")
	}
}

func returnBook(scanner *bufio.Scanner, library *services.Library) {
	fmt.Print("Enter book ID to return: ")
	scanner.Scan()
	bookID, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Enter member ID: ")
	scanner.Scan()
	memberID, _ := strconv.Atoi(scanner.Text())

	err := library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned.")
	}
}

func listAvailableBooks(library *services.Library) {
	books := library.ListAvailableBooks()
	fmt.Println("Available Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func listBorrowedBooks(scanner *bufio.Scanner, library *services.Library) {
	fmt.Print("Enter member ID: ")
	scanner.Scan()
	memberID, _ := strconv.Atoi(scanner.Text())

	books := library.ListBorrowedBooks(memberID)
	fmt.Println("Borrowed Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}
