package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

import _ "github.com/jinzhu/gorm/dialects/postgres"

// NewDataBase – returns a pointer to a database object
func NewDataBase() (*gorm.DB, error) {

	fmt.Println("Setting up new database connection")

	dbUsername := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbTable := os.Getenv("POSTGRES_TABLE")
	dbPort := os.Getenv("POSTGRES_PORT")

	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUsername, dbTable, dbPassword)
	fmt.Println(connectString)

	db, err := gorm.Open("postgres", connectString)

	if err != nil {
		return db, err
	}

	if err := db.DB().Ping(); err != nil {
		return db, err
	}

	return db, nil
}
