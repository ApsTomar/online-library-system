package data_store

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"github.com/library/envConfig"
	"github.com/library/migrations"
	"github.com/library/models"
)

type DataStore struct {
	Db *gorm.DB
}

type DbUtil interface {
	InsertData
	GetData
	BookIssue
	VerifyUser(models.LoginDetails) (*models.Account, error)
}

type InsertData interface {
	CreateUserAccount(models.Account) error
	CreateAuthor(author models.Author) error
	CreateBook(book models.Book) error
	CreateSubject(book models.Subject) error
}

type GetData interface {
	GetSubjects() (*[]models.Subject, error)
	GetAuthors() (*[]models.Author, error)
	GetBooks() (*[]models.Book, error)
	GetBooksByName(string) (*[]models.Book, error)
	GetBookByID(uint) (*models.Book, error)
	GetBooksByAuthor(uint) (*[]models.Book, error)
	GetBooksBySubject(uint) (*[]models.Book, error)
	GetAuthorsByName(string) (*[]models.Author, error)
	GetAuthorByID(uint) (*models.Author, error)
}

type BookIssue interface {
	GetHistory(string) (*[]models.BookHistory, error)
	GetCompleteHistory() (*[]models.BookHistory, error)
	CheckAvailability(uint) (bool, error)
	IssueBook(uint, uint) error
	ReturnBook(uint) error
}

func DbConnect(dbConfig *envConfig.Env) DbUtil {
	db, err := gorm.Open(dbConfig.SqlDialect, dbConfig.SqlUrl)
	if err != nil {
		glog.Fatalf("DB connection not established due to: %v", err)
	}
	err = migrations.InitMySQL(db)
	if err != nil {
		glog.Fatalf("error running migrations: %v", err)
	}
	return &DataStore{Db: db}
}
