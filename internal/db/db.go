package db

import (
	"fmt"
	"os"
	"todoApi/models"

	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	host     string
	port     string
	user     string
	password string
	dbname   string
)

var dbConnection *gorm.DB

// Gets the DB connection details from the .env details
func getDbConnectionDetails() {

	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname = os.Getenv("DB_NAME")
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")

}

// InitDB Creates the DB connection
func InitDB() {
	var err error

	getDbConnectionDetails()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}

	err = db.AutoMigrate(&models.TblTodo{})

	// Creates the table if it doesn't exist'
	if err != nil {
		log.Fatal("unable to migrate data", err)
	}

	fmt.Println("✅ Connected to MySQL")
}

func OpenDbConnection() error {
	var err error

	getDbConnectionDetails()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	dbConnection, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
		return err
	}

	return nil
}
