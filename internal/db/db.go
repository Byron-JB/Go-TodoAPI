package db

import (
	"fmt"
	"os"
	"todoApi/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	host     string
	port     string
	user     string
	password string
	dbname   string
	dsn      string
)

//var dbConnection *gorm.DB

type GormDatabase struct {
	dbConnection gorm.DB
}

// Gets the DB connection details from the .env details
func getDbConnectionDetails() {

	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname = os.Getenv("DB_NAME")
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

}

// Database interface
type Database interface {
	SaveTodosToDb(todos []models.TodoDto) ([]models.TodoDto, error)
	FetchTodosFromDb(skip int, take int) ([]models.TodoDto, error)
	DeleteTodoFromDb(id int) error
	UpdateTodosInDB(todos []models.TodoDto) ([]models.TodoDto, error)
}

// NewDatabase creates a new database instance
func NewDatabase() (*GormDatabase, error) {
	getDbConnectionDetails()

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = connection.AutoMigrate(&models.TblTodo{})
	if err != nil {
		return nil, err
	}
	return &GormDatabase{dbConnection: *connection}, nil
}
