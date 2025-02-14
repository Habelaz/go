package controllers

import (
	"fmt"
	"library-managment/models"
	"library-managment/services"
)

func LibraryCLI() {
	library := &services.Library{
		Books: make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addBook(library)
		case 2:
			removeBook(library)
		case 3:
			borrowBook(library)
		case 4:
			returnBook(library)
		case 5:
			listAvailableBooks(library)
		case 6:
			listBorrowedBooks(library)
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, try again.")
		}
	}

}

func addBook(library *services.Library){
	var book models.Book
	fmt.Print("Enter Book ID: ")
	fmt.Scanln(&book.ID)
	fmt.Print("Enter Book Title: ")
	fmt.Scanln(&book.Title)
	fmt.Print("Enter Book Author: ")
	fmt.Scanln(&book.Author)

	library.AddBook(book)
	fmt.Println("Book added successfully.")
}

func removeBook(library *services.Library){
	var bookID int
	fmt.Print("Enter Book ID: ")
	fmt.Scanln(&bookID)

	library.RemoveBook(bookID)
	fmt.Println("Book removed successfully.")
}

func borrowBook(library *services.Library){
	var bookID, memberID int
	fmt.Print("Enter Book ID: ")
	fmt.Scanln(&bookID)
	fmt.Print("Enter Member ID: ")
	fmt.Scanln(&memberID)

	err := library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}
}

func returnBook(library *services.Library){
	var bookID, memberID int
	fmt.Print("Enter Book ID: ")
	fmt.Scanln(&bookID)
	fmt.Print("Enter Member ID: ")
	fmt.Scanln(&memberID)

	err := library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book returned successfully.")
	}
}

func listAvailableBooks(library *services.Library){
	books := library.ListAvailableBooks()
	for _,book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func listBorrowedBooks(library *services.Library){
	var memberID int
	fmt.Print("Enter Member ID: ")
	fmt.Scanln(&memberID)

	books := library.ListBorrowedBooks(memberID)
	for _,book := range books{
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}
