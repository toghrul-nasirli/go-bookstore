package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/toghrul-nasirli/go-bookstore/pkg/models"
	"github.com/toghrul-nasirli/go-bookstore/pkg/utils"
	"net/http"
	"strconv"
)

func GetBooks(writer http.ResponseWriter, request *http.Request) {
	newBooks := models.GetBooks()
	response, _ := json.Marshal(newBooks)

	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(http.StatusOK)
	if _, err := writer.Write(response); err != nil {
		fmt.Println("Error while getting books")
	}
}

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing book ID")
	}

	book, _ := models.GetBookById(ID)
	response, _ := json.Marshal(book)

	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(http.StatusOK)
	if _, err := writer.Write(response); err != nil {
		fmt.Println("Error while getting book by id")
	}
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	book := models.Book{}
	utils.ParseBody(request, &book)
	book.CreateBook()
	response, _ := json.Marshal(book)

	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(http.StatusOK)
	if _, err := writer.Write(response); err != nil {
		fmt.Println("Error while creating book")
	}
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing book ID")
	}

	book := models.DeleteBook(ID)
	response, _ := json.Marshal(book)

	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(http.StatusOK)
	if _, err := writer.Write(response); err != nil {
		fmt.Println("Error while deleting book")
	}
}

func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	var editedBook = &models.Book{}
	utils.ParseBody(request, editedBook)

	vars := mux.Vars(request)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		fmt.Println("Error while parsing book ID")
	}

	book, db := models.GetBookById(ID)
	if editedBook.Name != "" {
		book.Name = editedBook.Name
	}
	if editedBook.Author != "" {
		book.Author = editedBook.Author
	}
	if editedBook.Publication != "" {
		book.Publication = editedBook.Publication
	}
	db.Save(&book)

	response, _ := json.Marshal(book)

	writer.Header().Set("Content-Type", "pkglication/json")
	writer.WriteHeader(http.StatusOK)
	if _, err := writer.Write(response); err != nil {
		fmt.Println("Error while updating book")
	}
}
