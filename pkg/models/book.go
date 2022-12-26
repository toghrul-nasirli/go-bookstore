package models

import (
	"github.com/jinzhu/gorm"
	"github.com/toghrul-nasirli/go-bookstore/pkg/config"
)

var (
	db *gorm.DB
)

type Book struct {
	gorm.Model
	Name        string `gorm:"json:name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func GetBooks() []Book {
	var Books []Book
	db.Find(&Books)

	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var Book Book
	db := db.Where("ID=?", Id).Find(&Book)

	return &Book, db
}

func (book *Book) CreateBook() *Book {
	db.NewRecord(book)
	db.Create(&book)

	return book
}

func DeleteBook(Id int64) Book {
	var Book Book
	db.Where("ID=?", Id).Delete(&Book)

	return Book
}
