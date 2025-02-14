package services

import "library-managment/models"
import "errors"
import "fmt"

type Library struct{
	Books map[int]models.Book
	Members map[int]models.Member
}

type libraryManager interface{
	AddBook(book models.Book)
	RemoveBook(bookID int) 
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}



func (lib *Library) AddBook(book models.Book){
	if lib.Books == nil{
		lib.Books = make(map[int]models.Book)
	}
	book.Status = "available"
	lib.Books[book.ID] = book
}

func (lib *Library) RemoveBook(bookID int) {
	delete(lib.Books, bookID)
}

func (lib *Library) BorrowBook(bookID int,memberID int) error{
	book, exist := lib.Books[bookID]	
	if !exist {
		return errors.New("book not found")
	}
	if book.Status == "borrowed" {
		return errors.New("book is already borrowed")
	}

	member,exist := lib.Members[memberID]
	if !exist {
		return errors.New("user not found")
	}

	book.Status = "borrowed"
	lib.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	lib.Members[memberID] = member

	return nil
}

func (lib *Library) ReturnBook(bookID int, memberID int) error {
	member,exist := lib.Members[memberID]
	if !exist {
		return errors.New("user not found")
	}

	found := false
	for i,book := range member.BorrowedBooks{
		if book.ID == bookID {
			found = true
			member.BorrowedBooks = append(member.BorrowedBooks[:i],member.BorrowedBooks[i+1:]...)
		}
	}
	if !found {
		return errors.New("book not borrowed by this member")
	}

	book := lib.Books[bookID]
	book.Status = "available"
	lib.Books[bookID] = book
	lib.Members[memberID] = member

	return nil
}

func (lib *Library) ListAvailableBooks() []models.Book{
	var availableBooks []models.Book
	for _,book := range lib.Books{
		if book.Status =="available"{
			availableBooks = append(availableBooks, book)
		}
	}

	return availableBooks
}

func (lib *Library) ListBorrowedBooks(memberID int)	[]models.Book{
	member,exist := lib.Members[memberID]
	if !exist {
		fmt.Println("user not found")
		return nil
	}
	return member.BorrowedBooks	
}